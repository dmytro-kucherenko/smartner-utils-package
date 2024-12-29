package adapter

import (
	"context"
	"database/sql"
	"fmt"
	"net/http"
	"os/signal"
	"syscall"
	"time"

	"github.com/dmytro-kucherenko/smartner-utils-package/pkg/log/types"
	"github.com/dmytro-kucherenko/smartner-utils-package/pkg/server/adapters/gin/middlewares"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func ConnectSQL(connection string) *sql.DB {
	db, err := sql.Open("postgres", connection)
	if err != nil {
		panic("Database connection error")
	}

	return db
}

func CreateRouter(port int, isProd bool, clientURL string) (*gin.Engine, *http.Server) {
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

func CreateRoutes(router *gin.Engine, logger types.Logger) *gin.RouterGroup {
	api := router.Group("/api/v1")
	api.Use(middlewares.Logger(), middlewares.Error())

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	logger.Info("Docs path: /swagger/index.html")

	return api
}

func ServeGracefully(server *http.Server, logger types.Logger, timeout time.Duration) {
	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			panic(fmt.Sprintln("Serving error:", err))
		}
	}()

	signalCtx, signalCancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer signalCancel()

	<-signalCtx.Done()

	logger.Warn("Shutting down the server")
	shutdownCtx, shutdownCancel := context.WithTimeout(context.Background(), timeout)
	defer shutdownCancel()

	if err := server.Shutdown(shutdownCtx); err != nil {
		logger.Error("Gracefully shutdown error:", err)
	}
}
