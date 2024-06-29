package ping

import (
	"net/http"

	"github.com/a-h/templ"
)

type pingApi struct {
	msg string
}

func NewRoute() pingApi {
	return pingApi{
		msg: "pong",
	}
}

func (h pingApi) Get() http.Handler {
	return templ.Handler(pingIndex(h.msg))
}
