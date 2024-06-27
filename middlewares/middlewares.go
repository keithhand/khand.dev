package middlewares

import (
	"fmt"
	"html"
	"net/http"
)

type logger interface {
	Debug(string, ...any)
}

type middlewares struct {
	logger logger
	mm     []func(logger, http.Handler) http.Handler
}

func New(lgr logger, mm ...func(logger, http.Handler) http.Handler) func(http.Handler) http.Handler {
	mwrs := middlewares{
		logger: lgr,
		mm:     mm,
	}
	return mwrs.addToMux
}

func (mwrs middlewares) addToMux(mux http.Handler) http.Handler {
	for _, m := range mwrs.mm {
		mux = m(mwrs.logger, mux)
	}
	return mux
}

func HttpPathLogs(log logger, h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Debug(fmt.Sprintf("got %s request", html.EscapeString(r.URL.Path)))
		h.ServeHTTP(w, r)
	})
}
