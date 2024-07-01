package main

import (
	"context"
	"fmt"
	"os"

	"khand.dev/khand.dev/config"
	"khand.dev/khand.dev/handlers"
	"khand.dev/khand.dev/json"
	"khand.dev/khand.dev/logs"
	"khand.dev/khand.dev/middlewares"
	"khand.dev/khand.dev/routes"
	"khand.dev/khand.dev/server"
)

func run(ctx context.Context, out *os.File, _ []string) error {
	logs := logs.New(out)
	cnfg := config.New(logs)
	json := json.New(logs)
	hdlr := handlers.New(logs)

	rts := routes.New(
		logs,
		hdlr.Ping(),
		hdlr.Index(),
		hdlr.GitHub(cnfg, json),
	)

	srv := server.NewHttp(ctx, logs, cnfg).
		WithRoutes(rts).
		WithMiddlewares(middlewares.New(
			logs,
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
		logs.Error(fmt.Errorf("main: %w", err).Error())
		os.Exit(1)
	}
}
