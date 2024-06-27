package main

import (
	"context"
	"fmt"
	"os"

	"khand.dev/khand.dev/logs"
	"khand.dev/khand.dev/server"
)

func run(ctx context.Context, out *os.File, _ []string) error {
	// logger := logs.NewLogger(out)
	srv := server.New(ctx)
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
