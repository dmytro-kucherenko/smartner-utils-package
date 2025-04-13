package adapter

import (
	"github.com/dmytro-kucherenko/smartner-utils-package/pkg/server"
	"github.com/gin-gonic/gin"
)

type Controller interface {
	Init(router *gin.RouterGroup)
}

type Module interface {
	server.Module
	Controllers() []Controller
}
