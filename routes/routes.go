package routes

import "net/http"

type logger interface {
	Debug(string, ...any)
	Warn(string, ...any)
	Error(string, ...any)
}

type ping interface {
	Get(http.ResponseWriter, *http.Request)
}

type gitHub interface {
	GetRepos(http.ResponseWriter, *http.Request)
}

type routes struct {
	logger logger
	ping   ping
	gitHub gitHub
}

func New(lgr logger, ping ping, gh gitHub) func(*http.ServeMux) {
	rts := routes{
		logger: lgr,
		ping:   ping,
		gitHub: gh,
	}
	return rts.addToMux
}

func (r routes) addToMux(mux *http.ServeMux) {
	r.logger.Debug("starting adding routes to http mux...")
	mux.HandleFunc("GET /ping", r.ping.Get)
	mux.HandleFunc("GET /projects", r.gitHub.GetRepos)
	r.logger.Debug("... finished adding routes to http mux")
}
