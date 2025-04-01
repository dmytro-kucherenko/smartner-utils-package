package adapter

import (
	"github.com/dmytro-kucherenko/smartner-utils-package/pkg/server"
	"github.com/gin-gonic/gin"
)

func Get[R any, P any](router *gin.RouterGroup, request server.Request[R, P], options *RequestConfig) {
	router.GET(options.Path, handle(request, options, false)...)
}

func Post[R any, P any](router *gin.RouterGroup, request server.Request[R, P], options *RequestConfig) {
	router.POST(options.Path, handle(request, options, true)...)
}

func Put[R any, P any](router *gin.RouterGroup, request server.Request[R, P], options *RequestConfig) {
	router.PUT(options.Path, handle(request, options, true)...)
}

func Delete[R any, P any](router *gin.RouterGroup, request server.Request[R, P], options *RequestConfig) {
	router.DELETE(options.Path, handle(request, options, false)...)
}
