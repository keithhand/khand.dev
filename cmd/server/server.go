package main

import (
	"context"
	"fmt"
	"os"

	"khand.dev/khand.dev/config"
	"khand.dev/khand.dev/handlers"
	"khand.dev/khand.dev/json"
	"khand.dev/khand.dev/log"
	"khand.dev/khand.dev/middlewares"
	"khand.dev/khand.dev/routes"
	"khand.dev/khand.dev/server"
)

func run(ctx context.Context, out *os.File, _ []string) error {
	log := log.New(out)
	cnfg := config.New(log)
	json := json.New(log)

	rts := routes.New(
		log,
		handlers.Ping(),
		handlers.Index(),
		handlers.GitHub(cnfg, json),
	)

	srv := server.NewHttp(ctx, log, cnfg).
		WithRoutes(rts).
		WithMiddlewares(middlewares.New(
			log,
			middlewares.HttpPathLogs,
		))

	if err := srv.Start(); err != nil {
		return fmt.Errorf("starting server: %w", err)
	}
	return nil
}

func main() {
	ctx := context.Background()
	if err := run(ctx, os.Stdout, os.Args); err != nil {
		log.Error(fmt.Errorf("main: %w", err).Error())
		os.Exit(1)
	}
}
