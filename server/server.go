package server

import (
	"context"
	"fmt"
	"net/http"

	"khand.dev/khand.dev/config"
	"khand.dev/khand.dev/logs"
	"khand.dev/khand.dev/routes"
)

type server struct {
	Context context.Context
	Config  *config.Config
	Server  *http.Server
}

func New(ctx context.Context) *server {
	cfg := config.New()
	return &server{
		Context: ctx,
		Server: &http.Server{
			Addr:    fmt.Sprintf(":%d", cfg.ServerPort),
			Handler: routes.NewHandler(cfg),
		},
	}
}

func (app *server) Start() error {
	logs.Info("starting server:", "address", app.Server.Addr)
	if err := app.Server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		return fmt.Errorf("listen and serve: %w", err)
	}
	return nil
}
