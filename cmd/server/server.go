package main

import (
	"fmt"
	"net/http"
	"os"

	"khand.dev/khand.dev/config"
	"khand.dev/khand.dev/routes"
)

func main() {
	hdl := routes.NewHandler()
	srv := http.Server{
		Addr:    fmt.Sprintf(":%d", config.ServerPort),
		Handler: hdl,
	}

	fmt.Printf("starting server at: %v\n", srv.Addr)
	if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		fmt.Printf("error starting server: %s\n", err)
		os.Exit(1)
	}
}
