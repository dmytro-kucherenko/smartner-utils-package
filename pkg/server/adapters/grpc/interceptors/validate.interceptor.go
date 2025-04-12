package interceptors

import (
	"context"

	"github.com/bufbuild/protovalidate-go"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/proto"
)

func ValidateUnary() func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	validator, _ := protovalidate.New()

	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		message := req.(proto.Message)
		if err := validator.Validate(message); err != nil {
			return nil, err
		}

		return handler(ctx, req)
	}
}
