package server

import (
	"context"
)

type Request[R any, P any] func(ctx context.Context, params P) (result R, err error)
