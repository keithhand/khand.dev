package routes

import "net/http"

type Logger interface {
	Debug(string, ...any)
	Warn(string, ...any)
	Error(string, ...any)
}

type Getter interface {
	Get() http.Handler
}

type routes struct {
	logger Logger
	ping   Getter
	index  Getter
	gitHub Getter
}

func New(lgr Logger, ping Getter, idx Getter, gh Getter) func(*http.ServeMux) {
	return routes{
		logger: lgr,
		ping:   ping,
		index:  idx,
		gitHub: gh,
	}.addToMux
}

func (r routes) addToMux(mux *http.ServeMux) {
	r.logger.Debug("starting adding routes to http mux...")

	mux.Handle("GET /ping", r.ping.Get())
	mux.Handle("GET /", r.index.Get())
	mux.Handle("GET /projects", r.gitHub.Get())

	r.logger.Debug("... finished adding routes to http mux")
}
