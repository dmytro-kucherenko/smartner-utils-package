package startup

import (
	"github.com/dmytro-kucherenko/smartner-utils-package/pkg/server"
)

func Init(options server.StartupOptions) {
	server.ServeGracefully(options.Server, options.Logger, options.ShutdownTimeout)
}
