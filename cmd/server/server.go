package main

import (
	"fmt"
	"net/http"
	"os"

	"khand.dev/khand.dev/config"
	"khand.dev/khand.dev/routes"
)

func main() {
	cfg := config.New()
	hdl := routes.NewHandler()
	srv := http.Server{
		Addr:    fmt.Sprintf(":%d", cfg.Port),
		Handler: hdl,
	}

	fmt.Printf("starting server at: %s\n", srv.Addr)
	if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		fmt.Printf("error starting server: %s\n", err)
		os.Exit(1)
	}
}
