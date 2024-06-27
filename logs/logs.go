package logs

import (
	"io"
	"log/slog"
	"os"
)

var (
	Info  func(string, ...any) = slog.Info
	Debug func(string, ...any) = slog.Debug
	Error func(string, ...any) = slog.Error
	Warn  func(string, ...any) = slog.Warn
)

func Fatal(err error) {
	slog.Error(err.Error())
	os.Exit(1)
}

func New(out io.Writer) *slog.Logger {
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
