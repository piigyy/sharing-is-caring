package server

import (
	"context"
	"errors"
	"fmt"
	"net"
	"net/http"
	"strconv"
	"time"

	"github.com/piigyy/sharing-is-caring/pkg/logger"
	"google.golang.org/grpc"
)

type Adapter func(http.Handler) http.Handler
type server struct {
	ip       string
	port     string
	listener net.Listener
}

func New(port string) (*server, error) {
	addr := fmt.Sprintf(":%s", port)
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		return nil, fmt.Errorf("failed to create listener on %s: %w", addr, err)
	}

	return &server{
		ip:       listener.Addr().(*net.TCPAddr).IP.String(),
		port:     strconv.Itoa(listener.Addr().(*net.TCPAddr).Port),
		listener: listener,
	}, nil
}

func (s *server) ServeHTTP(ctx context.Context, srv *http.Server) error {
	const caller = "server.ServeHTTP"

	// Spawn a goroutine that listens for context closure. When the context is
	// closed, the server is stopped.
	errChan := make(chan error, 1)
	go func() {
		<-ctx.Done()

		shutDownCtx, done := context.WithTimeout(context.Background(), 5*time.Second)
		defer done()

		logger.Info(ctx, caller, "server receive signal to shutdown")
		errChan <- srv.Shutdown(shutDownCtx)
	}()

	// Run the server. This will block until the provided context is closed.
	if err := srv.Serve(s.listener); err != nil && !errors.Is(err, http.ErrServerClosed) {
		logger.Error(ctx, caller, "failed to serve: %v", err)
		return fmt.Errorf("failed to serve: %w", err)
	}

	if err := <-errChan; err != nil {
		logger.Error(ctx, caller, "failed to shutdown: %v", err)
		return fmt.Errorf("failed to shutdown: %w", err)
	}

	return nil
}

func (s *server) ServeHTTPHandler(ctx context.Context, handler http.Handler) error {
	return s.ServeHTTP(ctx, &http.Server{
		Handler: handler,
	})
}

func (s *server) ServeGRPC(ctx context.Context, srv *grpc.Server) error {
	const caller = "server.ServeGRPC"

	// Spawn a goroutine that listens for context closure. When the context is
	// closed, the server is stopped.
	errChan := make(chan error, 1)
	go func() {
		<-ctx.Done()

		logger.Info(ctx, caller, "server.Server: context closed")
		logger.Info(ctx, caller, "server.Server: shutting down")
		srv.GracefulStop()
	}()

	if err := srv.Serve(s.listener); err != nil && !errors.Is(err, grpc.ErrServerStopped) {
		logger.Error(ctx, caller, "failed to serve: %v", err)
		return fmt.Errorf("failed to serve: %v", err)
	}

	logger.Info(ctx, caller, "server.Server: serving stopped")

	select {
	case err := <-errChan:
		return fmt.Errorf("failed to shutdown: %v", err)
	default:
		return nil
	}
}

func Adapt(handler http.Handler, adapters ...Adapter) http.Handler {
	for i := len(adapters); i > 0; i-- {
		handler = adapters[i-1](handler)
	}
	return handler
}
