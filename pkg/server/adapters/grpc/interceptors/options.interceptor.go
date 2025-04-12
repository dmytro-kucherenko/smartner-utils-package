package interceptors

import (
	"context"
	"errors"
	"fmt"

	"github.com/dmytro-kucherenko/smartner-utils-package/pkg/meta"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

func OptionsUnary() func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		md, ok := metadata.FromIncomingContext(ctx)
		if !ok {
			return nil, errors.New("metadata not found")
		}

		fmt.Println(md)

		ctx, err := meta.SetOptionsMetadataContext(ctx, md)
		if err != nil {
			return nil, err
		}

		return handler(ctx, req)
	}
}
