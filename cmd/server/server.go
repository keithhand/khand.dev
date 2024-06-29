package main

import (
	"context"
	"fmt"
	"os"

	"khand.dev/khand.dev/config"
	"khand.dev/khand.dev/json"
	"khand.dev/khand.dev/logs"
	"khand.dev/khand.dev/middlewares"
	"khand.dev/khand.dev/routes"
	"khand.dev/khand.dev/routes/github"
	"khand.dev/khand.dev/routes/ping"
	"khand.dev/khand.dev/server"
)

func run(ctx context.Context, out *os.File, _ []string) error {
	lgr := logs.New(out)
	cfg := config.New(lgr)
	jsn := json.New(lgr)
	mwr := middlewares.New(
		lgr,
		middlewares.HttpPathLogs,
	)
	rts := routes.New(
		lgr,
		ping.NewRoute(),
		github.NewRoute(lgr, jsn, cfg.GHProfile()),
	)
	srv := server.NewHttp(ctx, lgr, cfg.Port()).
		WithRoutes(rts).
		WithMiddlewares(mwr)
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
