package server

import (
	"context"
	"net/http"
	"time"

	"github.com/dmytro-kucherenko/smartner-utils-package/pkg/log/types"
)

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

type Session struct {
	UserID int `mapstructure:"userId" validate:"required,uuid4"`
}

type RequestMeta struct {
	Session *Session
}

type StartupOptions struct {
	Server          *http.Server
	Logger          types.Logger
	ShutdownTimeout time.Duration
}
