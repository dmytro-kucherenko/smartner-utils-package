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
	"github.com/dmytro-kucherenko/smartner-utils-package/pkg/server/errors"
)

func ConnectSQL(connection string, timeout time.Duration) (*sql.DB, error) {
	dbChan := make(chan *sql.DB, 1)
	errChan := make(chan error, 1)

	go func() {
		db, err := sql.Open("postgres", connection)
		if err != nil {
			errChan <- err

			return
		}

		err = db.Ping()
		if err != nil {
			errChan <- err

			return
		}

		dbChan <- db
	}()

	select {
	case db := <-dbChan:
		return db, nil
	case err := <-errChan:
		return nil, err
	case <-time.After(timeout):
		return nil, errors.NewHttpError(http.StatusGatewayTimeout, "database connection timeout reached")
	}
}

func ServeGracefully(server *http.Server, logger types.Logger, timeout time.Duration) error {
	errChan := make(chan error, 1)
	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			errChan <- fmt.Errorf("serving error: %v", err.Error())
		}
	}()

	signalCtx, signalCancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer signalCancel()

	select {
	case err := <-errChan:
		return err
	case <-signalCtx.Done():
		logger.Warn("shutting down the server")
		shutdownCtx, shutdownCancel := context.WithTimeout(context.Background(), timeout)
		defer shutdownCancel()

		if err := server.Shutdown(shutdownCtx); err != nil {
			logger.Error("gracefully shutdown error:", err)
		}
	}

	return nil
}
