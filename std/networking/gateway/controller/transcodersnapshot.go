// Copyright 2022 Namespace Labs Inc; All rights reserved.
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.

package main

import (
	"context"
	"encoding/base64"
	"fmt"
	"strings"
	"sync"
	"time"

	"go.uber.org/zap"
	"google.golang.org/protobuf/types/known/anypb"
	"google.golang.org/protobuf/types/known/durationpb"
	"namespacelabs.dev/foundation/framework/rpcerrors/multierr"
	"namespacelabs.dev/foundation/internal/fnerrors"

	accesslogv3 "github.com/envoyproxy/go-control-plane/envoy/config/accesslog/v3"
	cluster "github.com/envoyproxy/go-control-plane/envoy/config/cluster/v3"
	core "github.com/envoyproxy/go-control-plane/envoy/config/core/v3"
	endpoint "github.com/envoyproxy/go-control-plane/envoy/config/endpoint/v3"
	listener "github.com/envoyproxy/go-control-plane/envoy/config/listener/v3"
	route "github.com/envoyproxy/go-control-plane/envoy/config/route/v3"
	trace "github.com/envoyproxy/go-control-plane/envoy/config/trace/v3"
	filev3 "github.com/envoyproxy/go-control-plane/envoy/extensions/access_loggers/file/v3"
	gzipv3 "github.com/envoyproxy/go-control-plane/envoy/extensions/compression/gzip/compressor/v3"
	compressorv3 "github.com/envoyproxy/go-control-plane/envoy/extensions/filters/http/compressor/v3"
	grpcjsontranscoder "github.com/envoyproxy/go-control-plane/envoy/extensions/filters/http/grpc_json_transcoder/v3"
	routerfilter "github.com/envoyproxy/go-control-plane/envoy/extensions/filters/http/router/v3"
	hcm "github.com/envoyproxy/go-control-plane/envoy/extensions/filters/network/http_connection_manager/v3"
	tlstransport "github.com/envoyproxy/go-control-plane/envoy/extensions/transport_sockets/tls/v3"
	httpv3 "github.com/envoyproxy/go-control-plane/envoy/extensions/upstreams/http/v3"
	"github.com/envoyproxy/go-control-plane/pkg/cache/types"
	"github.com/envoyproxy/go-control-plane/pkg/cache/v3"
	"github.com/envoyproxy/go-control-plane/pkg/resource/v3"
	"github.com/envoyproxy/go-control-plane/pkg/wellknown"
	"github.com/golang/protobuf/ptypes/any"
	"google.golang.org/protobuf/proto"
	descriptorpb "google.golang.org/protobuf/types/descriptorpb"
)

const (
	ListenerName     = "listener_0"
	LocalRouteName   = "local_route"
	LocalVirtualHost = "local_virtual_host"
	StatPrefix       = "grpc_json"

	OtelClusterName = "otel-cluster"
)

type httpListenerConfig struct {
	name     string
	addrPort *AddressPort
}

type transcoderWithCluster struct {
	transcoder  *HttpGrpcTranscoder
	clusterName string
}

type TranscoderSnapshot struct {
	// Guards access to data below.
	mu sync.Mutex

	// Configuration of the http listener built from the registered transcoders.
	httpConfig httpListenerConfig

	// Envoy Node that we set the snapshot for.
	envoyNodeId string

	// Monotonically increasing counter of cache snapshot identifiers.
	snapshotId int
	cache      cache.SnapshotCache

	// Maps fully qualified proto service names to the corresponding HttpGrpcTranscoder.
	transcoders map[string]*HttpGrpcTranscoder

	// Default clusters that should be always created. Since each envoy snapshot overrides
	// the previous value, we need to keep a copy of the default bootstrapped clusters.
	defaultClusters []types.Resource

	otelCluster *cluster.Cluster
}

type SnapshotOptions struct {
	envoyNodeId string

	logger *zap.SugaredLogger

	xdsClusterName string
	xdsClusterAddr *AddressPort

	alsClusterName string
	alsClusterAddr *AddressPort

	otelEndpoint *AddressPort
}

type SnapshotOption func(*SnapshotOptions)

func WithEnvoyNodeId(envoyNodeId string) SnapshotOption {
	return func(o *SnapshotOptions) {
		o.envoyNodeId = envoyNodeId
	}
}

func WithLogger(logger *zap.SugaredLogger) SnapshotOption {
	return func(o *SnapshotOptions) {
		o.logger = logger
	}
}

func WithXdsCluster(xdsClusterName string, xdsClusterAddr *AddressPort) SnapshotOption {
	return func(o *SnapshotOptions) {
		o.xdsClusterName = xdsClusterName
		o.xdsClusterAddr = xdsClusterAddr
	}
}

func WithAlsCluster(alsClusterName string, alsClusterAddr *AddressPort) SnapshotOption {
	return func(o *SnapshotOptions) {
		o.alsClusterName = alsClusterName
		o.alsClusterAddr = alsClusterAddr
	}
}

func WithOtel(addrPort *AddressPort) SnapshotOption {
	return func(o *SnapshotOptions) {
		o.otelEndpoint = addrPort
	}
}

func NewTranscoderSnapshot(args ...SnapshotOption) (*TranscoderSnapshot, error) {
	opts := &SnapshotOptions{}
	for _, opt := range args {
		opt(opts)
	}

	var defaultClusters []types.Resource

	var errs []error

	if xdsCluster, err := makeCluster(opts.xdsClusterName, opts.xdsClusterAddr.addr, opts.xdsClusterAddr.port, false); err != nil {
		errs = append(errs, err)
	} else {
		defaultClusters = append(defaultClusters, xdsCluster)
	}

	if alsCluster, err := makeCluster(opts.alsClusterName, opts.alsClusterAddr.addr, opts.alsClusterAddr.port, false); err != nil {
		errs = append(errs, err)
	} else {
		defaultClusters = append(defaultClusters, alsCluster)
	}

	var otelCluster *cluster.Cluster
	if opts.otelEndpoint != nil {
		cluster, err := makeCluster(OtelClusterName, opts.otelEndpoint.addr, opts.otelEndpoint.port, false)
		if err != nil {
			errs = append(errs, err)
		} else {
			otelCluster = cluster
		}
	}

	if len(errs) > 0 {
		return nil, multierr.New(errs...)
	}

	cache := cache.NewSnapshotCache(false, cache.IDHash{}, opts.logger)
	return &TranscoderSnapshot{
		envoyNodeId:     opts.envoyNodeId,
		snapshotId:      1,
		cache:           cache,
		transcoders:     make(map[string]*HttpGrpcTranscoder),
		defaultClusters: defaultClusters,
		otelCluster:     otelCluster,
	}, nil
}

func (t *TranscoderSnapshot) CurrentSnapshotId() int {
	t.mu.Lock()
	defer t.mu.Unlock()
	return t.snapshotId
}

func (t *TranscoderSnapshot) RegisterHttpListener(listenerAddr string) error {
	t.mu.Lock()
	defer t.mu.Unlock()

	addrPort, err := ParseAddressPort(listenerAddr)
	if err != nil {
		return err
	}

	t.httpConfig = httpListenerConfig{ListenerName, addrPort}
	return nil
}

// AddTranscoder adds a new transformer.
func (t *TranscoderSnapshot) AddTranscoder(transcoder *HttpGrpcTranscoder) {
	t.mu.Lock()
	t.transcoders[transcoder.Spec.FullyQualifiedProtoServiceName] = transcoder
	t.mu.Unlock()
}

// DeleteTranscoder deletes a transcoder.
func (t *TranscoderSnapshot) DeleteTranscoder(transcoder *HttpGrpcTranscoder) {
	t.mu.Lock()
	delete(t.transcoders, transcoder.Spec.FullyQualifiedProtoServiceName)
	t.mu.Unlock()
}

// GenerateSnapshot generates a new envoy snapshot of all registered transcoders.
func (t *TranscoderSnapshot) GenerateSnapshot(ctx context.Context) error {
	t.mu.Lock()
	defer t.mu.Unlock()

	var transcoders []transcoderWithCluster

	var clusters []types.Resource
	clusters = append(clusters, t.defaultClusters...)

	var otel string
	if t.otelCluster != nil {
		clusters = append(clusters, t.otelCluster)
		otel = OtelClusterName
	}

	var errs []error
	for _, transcoder := range t.transcoders {
		clusterName := fmt.Sprintf("cluster-%s", strings.ReplaceAll(transcoder.Spec.FullyQualifiedProtoServiceName, ".", "-"))
		transcoders = append(transcoders, transcoderWithCluster{transcoder, clusterName})
		if cluster, err := makeCluster(clusterName, transcoder.Spec.ServiceAddress, transcoder.Spec.ServicePort, transcoder.Spec.BackendTLS); err != nil {
			errs = append(errs, err)
		} else {
			clusters = append(clusters, cluster)
		}
	}
	if len(errs) > 0 {
		return multierr.New(errs...)
	}

	httpListener, err := makeHTTPListener(t.httpConfig, transcoders, otel)
	if err != nil {
		return fnerrors.InternalError("failed to create the http listener: %w", err)
	}

	snapshot, err := cache.NewSnapshot(fmt.Sprintf("v.%d", t.snapshotId),
		map[resource.Type][]types.Resource{
			resource.ClusterType:  clusters,
			resource.ListenerType: {httpListener},
		},
	)
	if err != nil {
		return err
	}

	if err := snapshot.Consistent(); err != nil {
		return fnerrors.InternalError("failed to generate a consistent snapshot: %w", err)
	}

	if err := t.cache.SetSnapshot(ctx, t.envoyNodeId, snapshot); err != nil {
		return fnerrors.InternalError("failed to set the snapshot: %w", err)
	}

	// Increment the snapshot identifier after verifying everything is consistent.
	t.snapshotId++

	return nil
}

func makeCluster(clusterName string, socketAddress string, port uint32, backendTLS bool) (*cluster.Cluster, error) {
	httpopts := &httpv3.HttpProtocolOptions{
		UpstreamProtocolOptions: &httpv3.HttpProtocolOptions_ExplicitHttpConfig_{
			ExplicitHttpConfig: &httpv3.HttpProtocolOptions_ExplicitHttpConfig{
				ProtocolConfig: &httpv3.HttpProtocolOptions_ExplicitHttpConfig_Http2ProtocolOptions{},
			},
		},
	}
	httpoptsanypb, err := anypb.New(httpopts)
	if err != nil {
		return nil, fnerrors.InternalError("failed to serialize http options: %w", err)
	}

	cluster := &cluster.Cluster{
		Name:                 clusterName,
		ConnectTimeout:       durationpb.New(60 * time.Second),
		ClusterDiscoveryType: &cluster.Cluster_Type{Type: cluster.Cluster_LOGICAL_DNS},
		LbPolicy:             cluster.Cluster_ROUND_ROBIN,
		LoadAssignment:       makeEndpoint(clusterName, socketAddress, port),
		DnsLookupFamily:      cluster.Cluster_V4_ONLY,
		TypedExtensionProtocolOptions: map[string]*any.Any{
			"envoy.extensions.upstreams.http.v3.HttpProtocolOptions": httpoptsanypb,
		},
	}

	if backendTLS {
		// XXX security add server cert validation.
		tlsSocket := &tlstransport.UpstreamTlsContext{}

		tlsSocketAny, err := anypb.New(tlsSocket)
		if err != nil {
			return nil, fnerrors.InternalError("failed to serialize tls socket options: %w", err)
		}

		cluster.TransportSocket = &core.TransportSocket{
			Name: "envoy.transport_sockets.tls",
			ConfigType: &core.TransportSocket_TypedConfig{
				TypedConfig: tlsSocketAny,
			},
		}
	}

	return cluster, nil
}

func makeEndpoint(clusterName string, socketAddress string, port uint32) *endpoint.ClusterLoadAssignment {
	return &endpoint.ClusterLoadAssignment{
		ClusterName: clusterName,
		Endpoints: []*endpoint.LocalityLbEndpoints{{
			LbEndpoints: []*endpoint.LbEndpoint{{
				HostIdentifier: &endpoint.LbEndpoint_Endpoint{
					Endpoint: &endpoint.Endpoint{
						Address: &core.Address{
							Address: &core.Address_SocketAddress{
								SocketAddress: &core.SocketAddress{
									Protocol: core.SocketAddress_TCP,
									Address:  socketAddress,
									PortSpecifier: &core.SocketAddress_PortValue{
										PortValue: port,
									},
								},
							},
						},
					},
				},
			}},
		}},
	}
}

func decodeBase64FiledescriptorSet(encoded string) (*descriptorpb.FileDescriptorSet, error) {
	decodedContents, err := base64.StdEncoding.DecodeString(encoded)
	if err != nil {
		return nil, err
	}
	var fds descriptorpb.FileDescriptorSet
	if err := proto.Unmarshal(decodedContents, &fds); err != nil {
		return nil, err
	}
	if err != nil {
		return nil, err
	}
	return &fds, nil
}

// topoSort sorts FileDesciptorProtos such that imported files (dependencies) come first.
// If the file descriptor set is not topologically sorted and a dependency descriptor comes later,
// Envoy fails to build the descriptor pool correctly and throws an exception.
func topoSort(names []string, files map[string]*descriptorpb.FileDescriptorProto) []*descriptorpb.FileDescriptorProto {
	var result []*descriptorpb.FileDescriptorProto
	for _, name := range names {
		if file := files[name]; file != nil {
			result = append(result, topoSort(file.Dependency, files)...)
			result = append(result, file)
			delete(files, name)
		}
	}
	return result
}

func makeFiledescriptorSet(transcoders []transcoderWithCluster) (*descriptorpb.FileDescriptorSet, error) {
	files := map[string]*descriptorpb.FileDescriptorProto{}
	var names []string
	var errors []error

	for _, t := range transcoders {
		fileDescriptor, err := decodeBase64FiledescriptorSet(t.transcoder.Spec.EncodedProtoDescriptor)
		if err != nil {
			errors = append(errors, err)
		} else {
			for _, f := range fileDescriptor.File {
				name := f.GetName()
				if files[name] == nil {
					files[name] = f
					names = append(names, name)
				}
			}
		}
	}
	if len(errors) > 0 {
		return nil, multierr.New(errors...)
	}
	return &descriptorpb.FileDescriptorSet{
		File: topoSort(names, files),
	}, nil
}

func makeRoute(clusterName string, transcoderSpec HttpGrpcTranscoderSpec) *route.Route {
	return &route.Route{
		Match: &route.RouteMatch{
			PathSpecifier: &route.RouteMatch_Prefix{
				Prefix: "/" + transcoderSpec.FullyQualifiedProtoServiceName,
			},
		},
		Action: &route.Route_Route{
			Route: &route.RouteAction{
				ClusterSpecifier: &route.RouteAction_Cluster{
					Cluster: clusterName,
				},
				// Explicitly override the default upstream route timeout of 15s.
				// We set HttpConnectionManager.RequestTimeout to 5m below for unary requests.
				// https://www.envoyproxy.io/docs/envoy/latest/api-v3/config/route/v3/route_components.proto#config-route-v3-routeaction
				Timeout: durationpb.New(30 * time.Minute),
			},
		},
	}
}

func makeHTTPListener(httpConfig httpListenerConfig, transcoders []transcoderWithCluster, otelCluster string) (*listener.Listener, error) {
	var serviceNames []string
	var routes []*route.Route

	for _, t := range transcoders {
		serviceNames = append(serviceNames, t.transcoder.Spec.FullyQualifiedProtoServiceName)
		routes = append(routes, makeRoute(t.clusterName, t.transcoder.Spec))
	}

	fds, err := makeFiledescriptorSet(transcoders)
	if err != nil {
		return nil, fnerrors.InternalError("failed to created an aggregated FiledescriptorSet: %w", err)
	}
	bytes, err := proto.Marshal(fds)
	if err != nil {
		return nil, fnerrors.InternalError("failed to marshal the FiledescriptorSet: %w", err)
	}
	transcoderPb := &grpcjsontranscoder.GrpcJsonTranscoder{
		Services:    serviceNames,
		AutoMapping: true,
		DescriptorSet: &grpcjsontranscoder.GrpcJsonTranscoder_ProtoDescriptorBin{
			ProtoDescriptorBin: bytes,
		},
		PrintOptions: &grpcjsontranscoder.GrpcJsonTranscoder_PrintOptions{
			// Whether to add spaces, line breaks and indentation to make the JSON
			// output easy to read. Defaults to false.
			AddWhitespace: true,
			// Ensures that primitive fields with default values will be omitted in the JSON output.
			// For E.g., an int32 set to 0 or a string set to "".
			AlwaysPrintPrimitiveFields: false,
			// Ensures that enums are represented as strings in the JSON output.
			AlwaysPrintEnumsAsInts: false,
			// Whether to preserve proto field names. By default protobuf will
			// generate JSON field names using the ``json_name`` option, or lower camel case,
			// in that order. Setting this flag will preserve the original field names.
			PreserveProtoFieldNames: true,
		},
	}
	transcoderpbst, err := anypb.New(transcoderPb)
	if err != nil {
		return nil, fnerrors.BadInputError("failed to create the transcoder anypb: %w", err)
	}
	routerconfig, err := anypb.New(&routerfilter.Router{})
	if err != nil {
		return nil, fnerrors.BadInputError("failed to create the routerconfig anypb: %w", err)
	}
	fileAccessLog, err := anypb.New(&filev3.FileAccessLog{Path: "/dev/stdout"})
	if err != nil {
		return nil, fnerrors.BadInputError("failed to create fileaccesslog anypb: %w", err)
	}

	gzipConfig, err := anypb.New(&gzipv3.Gzip{})
	if err != nil {
		return nil, fnerrors.BadInputError("failed to create the gzipConfig anypb: %w", err)
	}
	compressorConfig, err := anypb.New(&compressorv3.Compressor{
		CompressorLibrary: &core.TypedExtensionConfig{
			Name:        "gzip",
			TypedConfig: gzipConfig,
		},
	})
	if err != nil {
		return nil, fnerrors.BadInputError("failed to create the compressorConfig anypb: %w", err)
	}

	// HTTP filter configuration
	manager := &hcm.HttpConnectionManager{
		CodecType:  hcm.HttpConnectionManager_AUTO,
		StatPrefix: StatPrefix,
		AccessLog: []*accesslogv3.AccessLog{{
			Name: wellknown.FileAccessLog,
			ConfigType: &accesslogv3.AccessLog_TypedConfig{
				TypedConfig: fileAccessLog,
			},
		}},
		StreamIdleTimeout: durationpb.New(30 * time.Minute),
		RequestTimeout:    durationpb.New(5 * time.Minute), // Does not apply to streaming requests.
		RouteSpecifier: &hcm.HttpConnectionManager_RouteConfig{
			RouteConfig: &route.RouteConfiguration{
				Name: LocalRouteName,
				VirtualHosts: []*route.VirtualHost{{
					Name:    LocalVirtualHost,
					Domains: []string{"*"},
					Routes:  routes,
				},
				}}},
		HttpFilters: []*hcm.HttpFilter{{
			Name: "envoy.filters.http.compressor",
			ConfigType: &hcm.HttpFilter_TypedConfig{
				TypedConfig: compressorConfig,
			},
		}, {
			Name: wellknown.GRPCJSONTranscoder,
			ConfigType: &hcm.HttpFilter_TypedConfig{
				TypedConfig: transcoderpbst,
			},
		}, {
			Name: wellknown.Router,
			ConfigType: &hcm.HttpFilter_TypedConfig{
				TypedConfig: routerconfig,
			},
		}},
	}

	if otelCluster != "" {
		openTelemetryConfig, err := anypb.New(&trace.OpenTelemetryConfig{
			GrpcService: &core.GrpcService{
				TargetSpecifier: &core.GrpcService_EnvoyGrpc_{
					EnvoyGrpc: &core.GrpcService_EnvoyGrpc{
						ClusterName: otelCluster,
					},
				},
				Timeout: durationpb.New(250 * time.Millisecond),
			},
			ServiceName: "envoy-grpc-transcoder",
		})
		if err != nil {
			return nil, fnerrors.BadInputError("failed to serialize OpenTelemetryConfig")
		}

		manager.Tracing = &hcm.HttpConnectionManager_Tracing{
			Provider: &trace.Tracing_Http{
				Name: "envoy.tracers.opentelemetry",
				ConfigType: &trace.Tracing_Http_TypedConfig{
					TypedConfig: openTelemetryConfig,
				},
			},
		}
	}

	pbst, err := anypb.New(manager)
	if err != nil {
		return nil, err
	}

	return &listener.Listener{
		Name: httpConfig.name,
		Address: &core.Address{
			Address: &core.Address_SocketAddress{
				SocketAddress: &core.SocketAddress{
					Protocol: core.SocketAddress_TCP,
					Address:  httpConfig.addrPort.addr,
					PortSpecifier: &core.SocketAddress_PortValue{
						PortValue: httpConfig.addrPort.port,
					},
				},
			},
		},
		FilterChains: []*listener.FilterChain{{
			Filters: []*listener.Filter{{
				Name: wellknown.HTTPConnectionManager,
				ConfigType: &listener.Filter_TypedConfig{
					TypedConfig: pbst,
				},
			}},
		}},
	}, nil
}
