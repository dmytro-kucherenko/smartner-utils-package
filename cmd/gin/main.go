package main

import (
	"time"

	"github.com/dmytro-kucherenko/smartner-utils-package/pkg/log"
	"github.com/dmytro-kucherenko/smartner-utils-package/pkg/server"
	adapter "github.com/dmytro-kucherenko/smartner-utils-package/pkg/server/adapters/gin"
)

const ShutdownTimeout time.Duration = 10 * time.Second

type Response struct {
	Success bool `json:"success"`
}

func route(options *server.RequestOptions[any, any, any]) (response Response, err error) {
	response.Success = true

	return
}

func main() {
	var meta server.RequestMeta
	logger := log.New("Init")

	router, httpServer := adapter.CreateRouter(8000, false, "mock")
	api := adapter.CreateRoutes(router, "/api", logger)

	config := adapter.NewConfig(meta)
	adapter.Post(api, route, config.MapRoute("/route", 200))

	server.ServeGracefully(httpServer, logger, ShutdownTimeout)
}
