package server

import (
	"context"
	"net/http"
	"time"

	loggerTypes "github.com/dmytro-kucherenko/smartner-utils-package/pkg/log/types"
	"github.com/dmytro-kucherenko/smartner-utils-package/pkg/types"
)

type RequestOptions[B any, P any, Q any] struct {
	Body    B
	Params  P
	Query   Q
	Ctx     context.Context
	Session Session
}

type Request[R any, B any, P any, Q any] func(data *RequestOptions[B, P, Q]) (result R, err error)

type Session struct {
	UserID types.ID `mapstructure:"userId" validate:"required,uuid4"`
}

type RequestMeta struct {
	Session *Session
}

type StartupOptions struct {
	Server          *http.Server
	Logger          loggerTypes.Logger
	ShutdownTimeout time.Duration
}
