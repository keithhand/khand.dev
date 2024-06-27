package routes

import (
	"fmt"
	"html"
	"net/http"

	"khand.dev/khand.dev/config"
	"khand.dev/khand.dev/logs"
)

func NewHandler(cfg *config.Config) http.Handler {
	mux := http.NewServeMux()
	addHandlers(mux, cfg)
	var handler http.Handler = mux
	handler = withLogs(handler)
	return handler
}

func addHandlers(mux *http.ServeMux, cfg *config.Config) {
	ping := NewPing()
	mux.HandleFunc("GET /ping", ping.Get)
	github := NewGitHubApi(cfg.GHProfile)
	mux.HandleFunc("GET /projects", github.GetProjects)
}

func withLogs(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logs.Debug(fmt.Sprintf("got %s request", html.EscapeString(r.URL.Path)))
		h.ServeHTTP(w, r)
	})
}
