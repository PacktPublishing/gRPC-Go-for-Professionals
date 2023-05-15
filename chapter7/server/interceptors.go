package main

import (
	"context"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

func validateAuthToken(ctx context.Context) error {
	md, ok := metadata.FromIncomingContext(ctx)

	if !ok {
		return status.Errorf(
			codes.DataLoss,
			"failed to get metadata",
		)
	}

	if t, ok := md["auth_token"]; ok {
		switch {
		case len(t) != 1:
			return status.Errorf(
				codes.InvalidArgument,
				"auth_token should contain only 1 value",
			)
		case t[0] != "authd":
			return status.Errorf(
				codes.Unauthenticated,
				"incorrect auth_token",
			)
		}
	} else {
		return status.Errorf(
			codes.Unauthenticated,
			"failed to get auth_token",
		)
	}

	return nil
}

func unaryAuthInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	if err := validateAuthToken(ctx); err != nil {
		return nil, err
	}

	return handler(ctx, req)
}

func streamAuthInterceptor(srv interface{}, ss grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
	if err := validateAuthToken(ss.Context()); err != nil {
		return err
	}

	return handler(srv, ss)
}

func unaryLogInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	log.Println(info.FullMethod, "called")
	return handler(ctx, req)
}

func streamLogInterceptor(srv interface{}, ss grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
	log.Println(info.FullMethod, "called")
	return handler(srv, ss)
}
