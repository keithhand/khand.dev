package routes

import (
	"fmt"
	"html"
	"net/http"
)

func NewHandler() http.Handler {
	mux := http.NewServeMux()
	addHandlers(mux)
	var handler http.Handler = mux
	handler = withLogs(handler)
	return handler
}

func addHandlers(mux *http.ServeMux) {
	ping := NewPingService()
	mux.HandleFunc("GET /ping", ping.Get)
	github := NewGitHubApiService()
	mux.HandleFunc("GET /projects", github.Projects.Get)
}

func withLogs(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("got %s request\n", html.EscapeString(r.URL.Path))
		h.ServeHTTP(w, r)
	})
}
