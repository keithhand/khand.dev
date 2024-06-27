package main

import (
	"log"

	"khand.dev/khand.dev/server"
)

func main() {
	srv := server.New()
	if err := srv.Start(); err != nil {
		log.Fatalf("fatal error: %s\n", err)
	}
}
