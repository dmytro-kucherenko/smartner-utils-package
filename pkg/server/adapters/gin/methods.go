package adapter

import (
	"github.com/dmytro-kucherenko/smartner-utils-package/pkg/server"
	"github.com/gin-gonic/gin"
)

func Config(path string, status int, middlewares []gin.HandlerFunc) *server.RequestConfig[gin.HandlerFunc] {
	return &server.RequestConfig[gin.HandlerFunc]{Path: path, Status: status, Middlewares: middlewares}
}

func Get[R any, B any, P any, Q any](router *gin.RouterGroup, options *server.RequestConfig[gin.HandlerFunc], request server.Request[R, B, P, Q]) {
	router.GET(options.Path, handle(options.Status, options.Middlewares, request, false)...)
}

func Post[R any, B any, P any, Q any](router *gin.RouterGroup, options *server.RequestConfig[gin.HandlerFunc], request server.Request[R, B, P, Q]) {
	router.POST(options.Path, handle(options.Status, options.Middlewares, request, true)...)
}

func Put[R any, B any, P any, Q any](router *gin.RouterGroup, options *server.RequestConfig[gin.HandlerFunc], request server.Request[R, B, P, Q]) {
	router.PUT(options.Path, handle(options.Status, options.Middlewares, request, true)...)
}

func Delete[R any, B any, P any, Q any](router *gin.RouterGroup, options *server.RequestConfig[gin.HandlerFunc], request server.Request[R, B, P, Q]) {
	router.DELETE(options.Path, handle(options.Status, options.Middlewares, request, false)...)
}
