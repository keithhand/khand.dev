package handlers

import (
	"net/http"

	"github.com/a-h/templ"
)

type Logger interface {
	Debug(string, ...any)
	Warn(string, ...any)
}

type Handler interface {
	View() templ.Component
}

type handler struct {
	Handler
}

func New(hdl Handler) *handler {
	return &handler{
		Handler: hdl,
	}
}

func (h handler) Get() http.Handler {
	return templ.Handler(h.Handler.View())
}
