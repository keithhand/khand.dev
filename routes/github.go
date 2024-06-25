package routes

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

var repos []gitHubRepo

type gitHubRepo struct {
	RepoName    string `json:"name"`
	Description string `json:"description"`
	Url         string `json:"url"`
}

type gitHubApiService struct {
	Projects projects
}

func NewGitHubApiService() gitHubApiService {
	return gitHubApiService{
		Projects: projects{},
	}
}

type projects struct{}

func (h projects) Get(w http.ResponseWriter, r *http.Request) {
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
