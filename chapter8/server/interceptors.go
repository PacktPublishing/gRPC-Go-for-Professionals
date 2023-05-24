package main

import (
	"context"
	"fmt"
	"log"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"

	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/logging"
)

func validateAuthToken(ctx context.Context) (context.Context, error) {
	md, ok := metadata.FromIncomingContext(ctx)

	if !ok {
		return nil, status.Errorf(
			codes.DataLoss,
			"failed to get metadata",
		)
	}

	if t, ok := md["auth_token"]; ok {
		switch {
		case len(t) != 1:
			return nil, status.Errorf(
				codes.InvalidArgument,
				"auth_token should contain only 1 value",
			)
		case t[0] != "authd":
			return nil, status.Errorf(
				codes.Unauthenticated,
				"incorrect auth_token",
			)
		}
	} else {
		return nil, status.Errorf(
			codes.Unauthenticated,
			"failed to get auth_token",
		)
	}

	return ctx, nil
}

const grpcService = 5 // "grpc.service"
const grpcMethod = 7  //"grpc.method"

func logCalls(l *log.Logger) logging.Logger {
	return logging.LoggerFunc(func(_ context.Context, lvl logging.Level, msg string, fields ...any) {
		// f := make(map[string]any, len(fields)/2)
		// i := logging.Fields(fields).Iterator()

		// for i.Next() {
		// 	k, v := i.At()
		// 	f[k] = v
		// }

		switch lvl {
		case logging.LevelDebug:
			msg = fmt.Sprintf("DEBUG :%v", msg)
		case logging.LevelInfo:
			msg = fmt.Sprintf("INFO :%v", msg)
		case logging.LevelWarn:
			msg = fmt.Sprintf("WARN :%v", msg)
		case logging.LevelError:
			msg = fmt.Sprintf("ERROR :%v", msg)
		default:
			panic(fmt.Sprintf("unknown level %v", lvl))
		}

		l.Println(msg, fields[grpcService], fields[grpcMethod])
		//l.Println(msg, f[grpcService], f[grpcMethod])
	})
}
