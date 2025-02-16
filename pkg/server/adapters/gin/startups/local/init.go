package startup

import (
	"github.com/dmytro-kucherenko/smartner-utils-package/pkg/server"
	adapter "github.com/dmytro-kucherenko/smartner-utils-package/pkg/server/adapters/gin"
	"github.com/gin-gonic/gin"
)

func Init(options *server.StartupOptions[gin.Engine]) {
	adapter.ServeGracefully(options.Server, options.Logger, options.ShutdownTimeout)
}
