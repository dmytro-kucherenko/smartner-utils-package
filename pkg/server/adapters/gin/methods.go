package adapter

import (
	"github.com/dmytro-kucherenko/smartner-utils-package/pkg/server"
	"github.com/gin-gonic/gin"
)

func Get[R any, B any, P any, Q any](router *gin.RouterGroup, request server.Request[R, B, P, Q], options *RequestConfig) {
	router.GET(options.Path, handle(request, options, false)...)
}

func Post[R any, B any, P any, Q any](router *gin.RouterGroup, request server.Request[R, B, P, Q], options *RequestConfig) {
	router.POST(options.Path, handle(request, options, true)...)
}

func Put[R any, B any, P any, Q any](router *gin.RouterGroup, request server.Request[R, B, P, Q], options *RequestConfig) {
	router.PUT(options.Path, handle(request, options, true)...)
}

func Delete[R any, B any, P any, Q any](router *gin.RouterGroup, request server.Request[R, B, P, Q], options *RequestConfig) {
	router.DELETE(options.Path, handle(request, options, false)...)
}
