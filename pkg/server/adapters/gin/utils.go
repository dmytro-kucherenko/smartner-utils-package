package adapter

import (
	"fmt"
	"net/http"

	"github.com/dmytro-kucherenko/smartner-utils-package/pkg/log/types"
	"github.com/dmytro-kucherenko/smartner-utils-package/pkg/server/adapters/gin/middlewares"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func CreateRouter(port uint16, isProd bool, clientURL string) (*gin.Engine, *http.Server) {
	if isProd {
		gin.SetMode(gin.ReleaseMode)
	}

	router := gin.New()
	server := &http.Server{
		Addr:    fmt.Sprintf(":%v", port),
		Handler: router.Handler(),
	}

	if isProd {
		router.SetTrustedProxies([]string{clientURL})
	}

	return router, server
}

func CreateRoutes(router *gin.Engine, prefix string, logger types.Logger) *gin.RouterGroup {
	api := router.Group(prefix)
	api.Use(middlewares.Logger(), middlewares.Error())

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	logger.Info("Docs path: /swagger/index.html")

	return api
}
