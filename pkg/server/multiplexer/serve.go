package multiplexer

import (
	"context"
	"net"
	"net/http"
	"os/signal"
	"syscall"
	"time"

	"github.com/dmytro-kucherenko/smartner-utils-package/pkg/log/types"
	"google.golang.org/grpc"
)

func ServeGracefully(
	serve func(net.Listener) error,
	stop func() error,
	listener net.Listener,
	timeout time.Duration,
	logger types.Logger,
) {
	done := make(chan error, 1)
	go func() {
		if err := serve(listener); err != nil && err != http.ErrServerClosed {
			logger.Error("serving error: %v", err.Error())
		}

		close(done)
	}()

	signalCtx, signalCancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer signalCancel()

	select {
	case <-done:
		return
	case <-signalCtx.Done():
		logger.Warn("stopping server")
		done := make(chan bool)

		go func() {
			err := stop()
			if err != nil {
				logger.Error("stopping error: %v", err.Error())
			}

			close(done)
		}()

		select {
		case <-done:
			return
		case <-time.After(timeout):
			logger.Warn("graceful stop timeout reached")
		}
	}
}

func ServeHTTPGracefully(server *http.Server, listener net.Listener, timeout time.Duration, logger types.Logger) {
	ServeGracefully(
		func(listener net.Listener) error { return server.Serve(listener) },
		func() error { return server.Shutdown(context.Background()) },
		listener,
		timeout,
		logger,
	)
}

func ServeGRPCGracefully(server *grpc.Server, listener net.Listener, timeout time.Duration, logger types.Logger) {
	ServeGracefully(
		func(listener net.Listener) error { return server.Serve(listener) },
		func() error { server.GracefulStop(); return nil },
		listener,
		timeout,
		logger,
	)
}
