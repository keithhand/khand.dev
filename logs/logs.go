package logs

import (
	"io"
	"log/slog"
	"os"
)

type Logger interface {
	Info(string, ...any)
	Debug(string, ...any)
	Warn(string, ...any)
	Error(string, ...any)
}

var (
	logger Logger               = newLogger(os.Stdout)
	Info   func(string, ...any) = logger.Info
	Debug  func(string, ...any) = logger.Debug
	Error  func(string, ...any) = logger.Error
	Warn   func(string, ...any) = logger.Warn
)

func Fatal(err error) {
	slog.Error(err.Error())
	os.Exit(1)
}

func newLogger(out io.Writer) Logger {
	opts := newOpts()
	handler := newHandler(out, opts)
	logger := slog.New(handler)
	slog.SetDefault(logger)
	return logger
}

func newOpts() *slog.HandlerOptions {
	return &slog.HandlerOptions{
		Level: slog.LevelDebug,
	}
}

func newHandler(out io.Writer, opts *slog.HandlerOptions) slog.Handler {
	return handlerMap(out, opts)["text"]
}

func handlerMap(out io.Writer, opts *slog.HandlerOptions) map[string]slog.Handler {
	return map[string]slog.Handler{
		"text": slog.NewTextHandler(out, opts),
		"json": slog.NewJSONHandler(out, opts),
	}
}
