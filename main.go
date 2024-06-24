package main

import (
	"fmt"
	"net/http"
	"os"

	"khand.dev/khand.dev/config"
	"khand.dev/khand.dev/routes"
)

func main() {
	env := config.NewConfig()
	handler := routes.NewServerHandler()
	srv := http.Server{
		Addr:    fmt.Sprintf(":%d", env.Port),
		Handler: handler,
	}

	fmt.Printf("starting server at: %s\n", srv.Addr)
	if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		fmt.Printf("error starting server: %s\n", err)
		os.Exit(1)
	}
}
