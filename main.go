package main

import (
	"errors"
	"fmt"
	"html"
	"io"
	"net/http"
	"os"
	"strconv"
)

const (
	portEnvVar  = "APP_PORT"
	portDefault = "8080"
)

func getPing(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("got %s request\n", html.EscapeString(r.URL.Path))
	io.WriteString(w, "pong")
}

type Environment struct {
	Port int
}

func NewEnvironment() Environment {
	portEnv := os.Getenv(portEnvVar)
	if portEnv == "" {
		portEnv = portDefault
	}

	port, err := strconv.Atoi(portEnv)
	if err != nil {
		panic(err)
	}

	return Environment{
		Port: port,
	}
}

func main() {
	env := NewEnvironment()

	http.HandleFunc("GET /ping", getPing)

	server := http.Server{
		Addr: fmt.Sprintf(":%d", env.Port),
	}
	fmt.Printf("starting server at: %s\n", server.Addr)
	if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		fmt.Printf("error starting server: %s\n", err)
		os.Exit(1)
	}
}
