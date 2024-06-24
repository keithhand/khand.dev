package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"html"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
)

const (
	portEnvVar  = "APP_PORT"
	portDefault = "8080"
)

func getPing(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("got %s request\n", html.EscapeString(r.URL.Path))
	io.WriteString(w, "pong")
}

var repos []GitHubRepo

type GitHubRepo struct {
	RepoName    string `json:"name"`
	Description string `json:"description"`
	Url         string `json:"url"`
}

func getProjects(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("got %s request\n", html.EscapeString(r.URL.Path))
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

type Environment struct {
	Port int
}

func NewEnvironment() Environment {
	portEnv := os.Getenv(portEnvVar)
	if portEnv == "" {
		portEnv = portDefault
	}

	port, err := strconv.Atoi(portEnv)
	if err != nil {
		panic(err)
	}

	return Environment{
		Port: port,
	}
}

func main() {
	env := NewEnvironment()

	http.HandleFunc("GET /ping", getPing)
	http.HandleFunc("GET /projects", getProjects)

	server := http.Server{
		Addr: fmt.Sprintf(":%d", env.Port),
	}
	fmt.Printf("starting server at: %s\n", server.Addr)
	if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		fmt.Printf("error starting server: %s\n", err)
		os.Exit(1)
	}
}
