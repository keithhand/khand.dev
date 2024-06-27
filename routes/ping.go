package routes

import (
	"io"
	"net/http"
)

type ping struct{}

func NewPing() ping {
	return ping{}
}

func (h ping) Get(w http.ResponseWriter, _ *http.Request) {
	io.WriteString(w, "pong")
}
