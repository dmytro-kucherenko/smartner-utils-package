package startup

import (
	"github.com/dmytro-kucherenko/smartner-utils-package/pkg/log"
	"github.com/dmytro-kucherenko/smartner-utils-package/pkg/server"
)

func Init(options server.StartupOptions) {
	logger := log.New("Init")
	err := server.ServeGracefully(options.Server, logger, options.ShutdownTimeout)
	if err != nil {
		panic(err.Error())
	}
}
