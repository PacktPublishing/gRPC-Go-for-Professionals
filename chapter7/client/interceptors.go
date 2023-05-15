package main

import (
	"context"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

const authTokenKey string = "auth_token"
const authTokenValue string = "authd"

func unaryAuthInterceptor(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
	ctx = metadata.AppendToOutgoingContext(ctx, authTokenKey, authTokenValue)
	err := invoker(ctx, method, req, reply, cc, opts...)

	return err
}

func streamAuthInterceptor(ctx context.Context, desc *grpc.StreamDesc, cc *grpc.ClientConn, method string, streamer grpc.Streamer, opts ...grpc.CallOption) (grpc.ClientStream, error) {
	ctx = metadata.AppendToOutgoingContext(ctx, authTokenKey, authTokenValue)
	s, err := streamer(ctx, desc, cc, method, opts...)

	if err != nil {
		return nil, err
	}

	return s, nil
}
