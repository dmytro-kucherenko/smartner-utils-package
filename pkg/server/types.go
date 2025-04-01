package server

import (
	"context"
	"net/http"
	"time"

	"github.com/dmytro-kucherenko/smartner-utils-package/pkg/types"
)

type RequestHeader struct {
	TimeZone []string `json:"TimeZone" mapstructure:"timeZone" validate:"omitempty,max=1,dive,timezone"`
}

type RequestOptions[P any] struct {
	Params   P
	Ctx      context.Context
	Session  Session
	TimeZone string
}

type Request[R any, P any] func(data *RequestOptions[P]) (result R, err error)

type Session struct {
	UserID types.ID `json:"userId" mapstructure:"userId" validate:"required,uuid4"`
}

type ParamsValidator interface {
	Make(data any) error
}

type RequestMeta struct {
	Session   *Session
	Validator ParamsValidator
}

type StartupOptions struct {
	Server          *http.Server
	ShutdownTimeout time.Duration
	OnlyConfig      bool
}
