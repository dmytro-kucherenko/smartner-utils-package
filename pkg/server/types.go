package server

import "context"

type RequestOptions[B any, P any, Q any] struct {
	Body   B
	Params P
	Query  Q
	Ctx    context.Context
}

type RequestConfig[M any] struct {
	Path        string
	Status      int
	Middlewares []M
}

type Request[R any, B any, P any, Q any] func(data *RequestOptions[B, P, Q]) (result R, err error)
