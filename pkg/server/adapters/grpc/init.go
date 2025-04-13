package adapter

import (
	"maps"

	"github.com/dmytro-kucherenko/smartner-utils-package/pkg/server"
	"google.golang.org/grpc"
)

func InitModules(server *grpc.Server, modules ...server.Module) {
	for _, module := range modules {
		InitModules(server, module.Modules()...)

		moduleCurrent, ok := module.(Module)
		if ok {
			for _, controller := range moduleCurrent.Callers() {
				controller.Init(server)
			}
		}
	}
}

func GetConfig(modules ...server.Module) CallerConfig {
	config := make(CallerConfig)

	for _, module := range modules {
		GetConfig(module.Modules()...)

		moduleCurrent, ok := module.(Module)
		if ok {
			for _, caller := range moduleCurrent.Callers() {
				callerConfig := caller.Config()
				maps.Copy(config, callerConfig)
			}
		}
	}

	return config
}
