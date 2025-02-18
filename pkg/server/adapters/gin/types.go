package adapter

import (
	"github.com/dmytro-kucherenko/smartner-utils-package/pkg/server"
	"github.com/gin-gonic/gin"
)

type StartupOptions struct {
	server.StartupOptions
	Router *gin.Engine
}
