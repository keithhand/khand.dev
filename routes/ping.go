package routes

import (
	"io"
	"net/http"
)

type pingApi struct {
	msg string
}

func NewPing() pingApi {
	return pingApi{
		msg: "pong",
	}
}

func (h pingApi) Get(w http.ResponseWriter, _ *http.Request) {
	io.WriteString(w, h.msg)
}
