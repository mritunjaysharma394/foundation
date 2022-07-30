// Copyright 2022 Namespace Labs Inc; All rights reserved.
// Licensed under the EARLY ACCESS SOFTWARE LICENSE AGREEMENT
// available at http://github.com/namespacelabs/foundation

package fnapi

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"
	"strconv"

	"github.com/spf13/pflag"
	spb "google.golang.org/genproto/googleapis/rpc/status"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"namespacelabs.dev/foundation/internal/fnerrors"
	"namespacelabs.dev/foundation/workspace/tasks"
)

var EndpointAddress = "https://api.namespacelabs.net"

func SetupFlags(flags *pflag.FlagSet) {
	flags.StringVar(&EndpointAddress, "fnapi_endpoint", EndpointAddress, "The fnapi endpoint address.")
	_ = flags.MarkHidden("fnapi_endpoint")
}

func callProdAPI(ctx context.Context, method string, req interface{}, handle func(dec *json.Decoder) error) error {
	return tasks.Action("fnapi.call").LogLevel(2).IncludesPrivateData().Arg("endpoint", EndpointAddress).Arg("method", method).Arg("request", req).Run(ctx, func(ctx context.Context) error {
		return CallAPI(ctx, EndpointAddress, method, req, handle)
	})
}

func CallAPI(ctx context.Context, endpoint string, method string, req interface{}, handle func(dec *json.Decoder) error) error {
	return CallAPIRaw(ctx, endpoint, method, req, func(body io.Reader) error {
		return handle(json.NewDecoder(body))
	})
}

func CallAPIRaw(ctx context.Context, endpoint string, method string, req interface{}, handle func(body io.Reader) error) error {
	reqBytes, err := json.Marshal(req)
	if err != nil {
		return fnerrors.InvocationError("failed to marshal request: %w", err)
	}

	httpReq, err := http.NewRequestWithContext(ctx, http.MethodPost, endpoint+"/"+method, bytes.NewReader(reqBytes))
	if err != nil {
		return fnerrors.InvocationError("failed to construct request: %w", err)
	}

	c := &http.Client{}
	response, err := c.Do(httpReq)
	if err != nil {
		return fnerrors.InvocationError("failed to perform invocation: %w", err)
	}

	defer response.Body.Close()

	if response.StatusCode == http.StatusOK {
		return handle(response.Body)
	}

	st := &spb.Status{}
	dec := json.NewDecoder(response.Body)
	if err := dec.Decode(st); err == nil {
		if st.Code == int32(codes.Unauthenticated) {
			return ErrRelogin
		}

		return status.ErrorProto(st)
	}

	grpcMessage := response.Header[http.CanonicalHeaderKey("grpc-message")]
	grpcStatus := response.Header[http.CanonicalHeaderKey("grpc-status")]

	if len(grpcMessage) > 0 && len(grpcStatus) > 0 {
		intVar, err := strconv.Atoi(grpcStatus[0])
		if err == nil {
			st.Code = int32(intVar)
			st.Message = grpcMessage[0]
			return status.ErrorProto(st)
		}
	}

	switch response.StatusCode {
	case http.StatusInternalServerError:
		return fnerrors.InvocationError("internal server error, and wasn't able to parse error response")
	case http.StatusForbidden:
		return fnerrors.InvocationError("forbidden")
	case http.StatusUnauthorized:
		return ErrRelogin
	default:
		return fnerrors.InvocationError("unexpected %d error reaching %q: %s", response.StatusCode, endpoint, response.Status)
	}
}
