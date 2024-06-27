package server

import (
	"fmt"
	"log"
	"net/http"

	"khand.dev/khand.dev/config"
	"khand.dev/khand.dev/routes"
)

type Server struct {
	*http.Server
}

func New() Server {
	return Server{
		&http.Server{
			Addr:    fmt.Sprintf(":%d", config.ServerPort),
			Handler: routes.NewHandler(),
		},
	}
}

func (srv Server) Start() error {
	log.Printf("starting server at: %v\n", srv.Addr)
	if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		return fmt.Errorf("starting server: %w", err)
	}
	return nil
}
