package multiplexer

import (
	"fmt"
	"net"
	"net/http"
	"time"

	"github.com/dmytro-kucherenko/smartner-utils-package/pkg/log"
	"github.com/dmytro-kucherenko/smartner-utils-package/pkg/log/types"
	"github.com/soheilhy/cmux"
	"google.golang.org/grpc"
)

type Service struct {
	server  cmux.CMux
	port    uint16
	timeout time.Duration
	logger  types.Logger
}

func NewService(port uint16) (*Service, error) {
	listener, err := net.Listen("tcp", fmt.Sprintf(":%v", port))
	if err != nil {
		return nil, err
	}

	server := cmux.New(listener)

	return &Service{server, port, 10 * time.Second, log.New("multiplexer")}, nil
}

func (multiplexer *Service) Port() uint16 {
	return multiplexer.port
}

func (multiplexer *Service) Timeout() time.Duration {
	return multiplexer.timeout
}

func (multiplexer *Service) SetTimeout(timeout time.Duration) {
	multiplexer.timeout = timeout
}

func (multiplexer *Service) WithGRPC(server *grpc.Server) *Service {
	listener := multiplexer.server.MatchWithWriters(cmux.HTTP2MatchHeaderFieldSendSettings("content-type", "application/grpc"))
	go ServeGRPCGracefully(server, listener, multiplexer.timeout, multiplexer.logger)

	return multiplexer
}

func (multiplexer *Service) WithHTTP(server *http.Server) *Service {
	listener := multiplexer.server.Match(cmux.HTTP1Fast())
	go ServeHTTPGracefully(server, listener, multiplexer.timeout, multiplexer.logger)

	return multiplexer
}

func (multiplexer *Service) Serve() error {
	return multiplexer.server.Serve()
}
