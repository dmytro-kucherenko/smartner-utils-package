package server

import (
	"context"
	"database/sql"
	"fmt"
	"net/http"
	"os/signal"
	"syscall"
	"time"

	"github.com/dmytro-kucherenko/smartner-utils-package/pkg/log/types"
)

func ConnectSQL(connection string) *sql.DB {
	db, err := sql.Open("postgres", connection)
	if err != nil {
		panic("Database connection error")
	}

	return db
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
