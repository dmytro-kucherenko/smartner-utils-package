package main

import (
	"time"

	"github.com/dmytro-kucherenko/smartner-utils-package/pkg/log"
	schema "github.com/dmytro-kucherenko/smartner-utils-package/pkg/schema/adapters/playground"
	"github.com/dmytro-kucherenko/smartner-utils-package/pkg/server"
	adapter "github.com/dmytro-kucherenko/smartner-utils-package/pkg/server/adapters/gin"
)

const ShutdownTimeout time.Duration = 10 * time.Second

type Params struct {
	Field string `json:"field" form:"field" uri:"field" validate:"required,min=3"`
}

type Response struct {
	Success bool   `json:"success"`
	Result  string `json:"result"`
}

func route(options *server.RequestOptions[Params]) (response Response, err error) {
	response.Success = true
	response.Result = options.Params.Field

	return
}

func main() {
	validator, err := schema.NewParamsValidator()
	if err != nil {
		panic(err.Error())
	}

	meta := server.RequestMeta{Validator: validator}
	logger := log.New("Init")

	router, httpServer := adapter.CreateRouter(8000, false, "mock")
	api := adapter.CreateRoutes(router, "/api", logger)

	config := adapter.NewConfig(meta)
	adapter.Post(api, route, config.MapRoute("/route/:field", 200))

	server.ServeGracefully(httpServer, logger, ShutdownTimeout)
}
