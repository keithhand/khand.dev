package log

import (
	"io"
	"log/slog"
	"os"
	"strings"
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

type logger struct {
	*slog.Logger
}

func New(out io.Writer) *logger {
	opts := newOpts(loggerLevel())
	handler := newHandler(out, opts)
	log := slog.New(handler)
	slog.SetDefault(log)
	return &logger{
		Logger: log,
	}
}

func loggerLevel() slog.Level {
	switch strings.ToUpper(os.Getenv("LOG_LEVEL")) {
	case "ERROR":
		return slog.LevelError
	case "WARN":
		return slog.LevelWarn
	case "INFO":
		return slog.LevelInfo
	case "DEBUG":
		return slog.LevelDebug
	default:
		return slog.LevelInfo
	}
}

func newOpts(lvl slog.Level) *slog.HandlerOptions {
	return &slog.HandlerOptions{
		Level: lvl,
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
