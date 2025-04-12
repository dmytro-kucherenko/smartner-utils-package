package adapter

import (
	"net/http"

	"github.com/dmytro-kucherenko/smartner-utils-package/pkg/log/types"
	"github.com/dmytro-kucherenko/smartner-utils-package/pkg/server/adapters/gin/interceptors"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func CreateRouter(isProd bool, clientURL string) (*gin.Engine, *http.Server) {
	if isProd {
		gin.SetMode(gin.ReleaseMode)
	}

	router := gin.New()
	server := &http.Server{Handler: router.Handler()}

	if isProd {
		router.SetTrustedProxies([]string{clientURL})
	}

	return router, server
}

func CreateRoutes(router *gin.Engine, prefix string, logger types.Logger) *gin.RouterGroup {
	api := router.Group(prefix)
	api.Use(interceptors.Logger(), interceptors.Error(), interceptors.Options())

	api.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	logger.Info("docs path: /swagger/index.html")

	return api
}
