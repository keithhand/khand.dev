package routes

import (
	"encoding/json"
	"fmt"
	"html"
	"io"
	"log"
	"net/http"
)

func NewServerHandler() http.Handler {
	mux := http.NewServeMux()
	addRoutes(mux)
	var handler http.Handler = mux
	handler = withLogs(handler)
	return handler
}

func addRoutes(mux *http.ServeMux) {
	mux.HandleFunc("GET /ping", getPing)
	mux.HandleFunc("GET /projects", getProjects)
}

func withLogs(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("got %s request\n", html.EscapeString(r.URL.Path))
		h.ServeHTTP(w, r)
	})
}

func getPing(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "pong")
}

var repos []gitHubRepo

type gitHubRepo struct {
	RepoName    string `json:"name"`
	Description string `json:"description"`
	Url         string `json:"url"`
}

func getProjects(w http.ResponseWriter, r *http.Request) {
	defer func() {
		for i := range repos {
			io.WriteString(w, fmt.Sprintf("repo: %s\n", repos[i]))
		}
	}()

	if repos != nil {
		return
	}

	user := "keithhand"
	repoApi := fmt.Sprintf("https://api.github.com/users/%s/repos", user)
	resp, err := http.Get(repoApi)
	if err != nil {
		fmt.Printf("error getting repo information: %s\n", err)
	}

	if resp.Body != nil {
		defer resp.Body.Close()
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	if err = json.Unmarshal(body, &repos); err != nil {
		log.Fatalln(err)
	}
}
