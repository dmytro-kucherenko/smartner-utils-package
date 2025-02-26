package adapter

import (
	"github.com/dmytro-kucherenko/smartner-utils-package/pkg/server"
	"github.com/gin-gonic/gin"
)

type RequestConfig = server.RequestConfig[gin.HandlerFunc]

func NewConfig(meta server.RequestMeta) *RequestConfig {
	return server.NewConfig[gin.HandlerFunc](meta)
}
