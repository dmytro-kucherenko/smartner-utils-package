package interceptors

import (
	"context"
	"errors"

	"github.com/dmytro-kucherenko/smartner-utils-package/pkg/meta"
	adapter "github.com/dmytro-kucherenko/smartner-utils-package/pkg/server/adapters/grpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

func ConfigUnary(callers ...adapter.Caller) func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	config := make(adapter.CallerConfig)
	for _, caller := range callers {
		callerConfig := caller.Config()
		for method, methodConfig := range callerConfig {
			config[method] = methodConfig
		}
	}

	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		config := config[info.FullMethod]

		if !config.Public {
			md, ok := metadata.FromIncomingContext(ctx)
			if !ok {
				return nil, errors.New("metadata not found")
			}

			var err error
			ctx, err = meta.SetSessionMetadataContext(ctx, md)
			if err != nil {
				return nil, err
			}
		}

		return handler(ctx, req)
	}
}
