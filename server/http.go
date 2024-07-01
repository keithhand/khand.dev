package server

import (
	"context"
	"fmt"
	"net/http"
)

type Logger interface {
	Debug(string, ...any)
	Error(string, ...any)
	Warn(string, ...any)
	Info(string, ...any)
}

type Config interface {
	Port() int
}

type HttpServer struct {
	server  *http.Server
	context context.Context
	logger  Logger
}

func NewHttp(ctx context.Context, lgr Logger, cfg Config) *HttpServer {
	return &HttpServer{
		logger:  lgr,
		context: ctx,
		server: &http.Server{
			Addr:    fmt.Sprintf(":%d", cfg.Port()),
			Handler: http.NewServeMux(),
		},
	}
}

func (srv *HttpServer) WithRoutes(r func(*http.ServeMux)) *HttpServer {
	r(srv.server.Handler.(*http.ServeMux))
	return srv
}

func (srv *HttpServer) WithMiddlewares(m func(http.Handler) http.Handler) *HttpServer {
	srv.server.Handler = m(srv.server.Handler)
	return srv
}

func (srv *HttpServer) Start() error {
	srv.logger.Info("starting server:", "address", srv.server.Addr)
	if err := srv.server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		return fmt.Errorf("listen and serve: %w", err)
	}
	return nil
}
