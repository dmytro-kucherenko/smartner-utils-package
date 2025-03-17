package startup

import (
	"github.com/dmytro-kucherenko/smartner-utils-package/pkg/log"
	"github.com/dmytro-kucherenko/smartner-utils-package/pkg/log/types"
	"github.com/dmytro-kucherenko/smartner-utils-package/pkg/server"
	"github.com/google/uuid"
)

func InitWithMeta(create func(logger types.Logger, meta server.RequestMeta) (server.StartupOptions, error), meta server.RequestMeta) {
	logger := log.New("Init")
	options, err := create(logger, meta)
	if err != nil {
		panic(err.Error())
	}

	if options.OnlyConfig {
		logger.Info("config was checked")

		return
	}

	err = server.ServeGracefully(options.Server, logger, options.ShutdownTimeout)
	if err != nil {
		panic(err.Error())
	}
}

func Init(create func(logger types.Logger, meta server.RequestMeta) (server.StartupOptions, error)) {
	id, _ := uuid.Parse("451f4f07-5140-456f-9ffc-4751a808f45f")
	meta := server.RequestMeta{
		Session: &server.Session{
			UserID: id,
		},
	}

	InitWithMeta(create, meta)
}
