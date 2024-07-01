package handlers

import (
	"context"
	"io"
	"net/http"

	"github.com/a-h/templ"
)

type Logger interface {
	Debug(string, ...any)
	Warn(string, ...any)
}

type Viewer interface {
	Render(context.Context, io.Writer) error
}

type Handler interface {
	Get() http.Handler
	Viewer() Viewer
	setLogger(Logger) Handler
}

type handler struct {
	Handler
	logger Logger
	viewer Viewer
}

func New(lgr Logger) *handler {
	return &handler{
		logger: lgr,
	}
}

func (h handler) Get() http.Handler {
	return templ.Handler(h.Viewer())
}

func (h *handler) setLogger(lgr Logger) Handler {
	h.logger = lgr
	return h
}

func (h handler) Viewer() Viewer {
	return h.viewer
}
