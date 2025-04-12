package adapter

import (
	"context"
	"encoding/json"

	"google.golang.org/grpc"
)

func ParseMessage[T any](message any) (value T, err error) {
	data, err := json.Marshal(message)
	if err != nil {
		return
	}

	err = json.Unmarshal(data, &value)
	if err != nil {
		return
	}

	return
}

func HandleCall[R, P any, RM, PM any](
	handler func(context.Context, *PM, ...grpc.CallOption) (*RM, error),
	ctx context.Context,
	params P,
	result R,
	options ...grpc.CallOption,
) (R, error) {
	req, err := ParseMessage[PM](params)
	if err != nil {
		return result, err
	}

	response, err := handler(ctx, &req, options...)
	if err != nil {
		return result, err
	}

	result, err = ParseMessage[R](response)

	return result, err
}

func HandleProcedure[R, P any, RM, PM any](
	handler func(context.Context, P) (R, error),
	ctx context.Context,
	request *PM,
	response *RM,
	options ...grpc.CallOption,
) (resp *RM, err error) {
	params, err := ParseMessage[P](request)
	if err != nil {
		return response, err
	}

	result, err := handler(ctx, params)
	if err != nil {
		return response, err
	}

	value, err := ParseMessage[RM](result)
	response = &value

	return response, err
}
