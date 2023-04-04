// Copyright 2022 Namespace Labs Inc; All rights reserved.
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.

package api

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"net/url"
	"time"

	"github.com/gorilla/websocket"
	"github.com/jpillora/chisel/share/cnet"
	"namespacelabs.dev/foundation/internal/fnapi"
)

func DialPort(ctx context.Context, cluster *KubernetesCluster, targetPort int) (net.Conn, error) {
	token, err := fnapi.FetchTenantToken(ctx)
	if err != nil {
		return nil, err
	}

	return DialPortWithToken(ctx, token, cluster, targetPort)
}

func DialEndpoint(ctx context.Context, endpoint string) (net.Conn, error) {
	token, err := fnapi.FetchTenantToken(ctx)
	if err != nil {
		return nil, err
	}

	return DialEndpointWithToken(ctx, token, endpoint)
}

func DialPortWithToken(ctx context.Context, token fnapi.Token, cluster *KubernetesCluster, targetPort int) (net.Conn, error) {
	return DialEndpointWithToken(ctx, token, fmt.Sprintf("wss://gate.%s/%s/%d", cluster.IngressDomain, cluster.ClusterId, targetPort))
}

func DialHostedServiceWithToken(ctx context.Context, token fnapi.Token, cluster *KubernetesCluster, serviceName string, vars url.Values) (net.Conn, error) {
	u := url.URL{
		Scheme:   "wss",
		Host:     fmt.Sprintf("gate.%s", cluster.IngressDomain),
		Path:     fmt.Sprintf("/%s/hsvc.%s", cluster.ClusterId, serviceName),
		RawQuery: vars.Encode(),
	}

	return DialEndpointWithToken(ctx, token, u.String())
}

func DialEndpointWithToken(ctx context.Context, token fnapi.Token, endpoint string) (net.Conn, error) {
	d := websocket.Dialer{
		HandshakeTimeout: 15 * time.Second,
	}

	hdrs := http.Header{}
	hdrs.Add("Authorization", fnapi.BearerToken(token))

	wsConn, _, err := d.DialContext(ctx, endpoint, hdrs)
	if err != nil {
		return nil, err
	}

	return cnet.NewWebSocketConn(wsConn), nil
}
