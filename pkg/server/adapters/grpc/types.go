package adapter

import (
	"github.com/dmytro-kucherenko/smartner-utils-package/pkg/server"
	"google.golang.org/grpc"
)

type MethodConfig struct {
	Public bool
}

type CallerConfig map[string]MethodConfig

type Caller interface {
	Init(server *grpc.Server)
	Config() CallerConfig
}

type Module interface {
	server.Module
	Callers() []Caller
}
