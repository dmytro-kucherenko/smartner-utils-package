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
	UserID int `json:"userId" validate:"required"`
}

type RequestMeta struct {
	Session *Session `json:"session"`
}

type StartupOptions[R any] struct {
	Router          *R
	Server          *http.Server
	Logger          types.Logger
	ShutdownTimeout time.Duration
}
