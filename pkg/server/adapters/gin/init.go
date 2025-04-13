package adapter

import (
	"github.com/dmytro-kucherenko/smartner-utils-package/pkg/server"
	"github.com/gin-gonic/gin"
)

func InitModules(router *gin.RouterGroup, modules ...server.Module) {
	for _, module := range modules {
		InitModules(router, module.Modules()...)

		moduleCurrent, ok := module.(Module)
		if ok {
			for _, controller := range moduleCurrent.Controllers() {
				controller.Init(router)
			}
		}
	}
}
