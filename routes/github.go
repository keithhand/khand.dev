package routes

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"khand.dev/khand.dev/logs"
)

type gitHubApi struct {
	Profile string
}

func NewGitHubApi(profile string) gitHubApi {
	return gitHubApi{
		Profile: profile,
	}
}

var repos []gitHubRepo

type gitHubRepo struct {
	RepoName    string `json:"name"`
	Description string `json:"description"`
	Url         string `json:"url"`
}

func (api gitHubApi) GetProjects(w http.ResponseWriter, r *http.Request) {
	defer func() {
		for i := range repos {
			io.WriteString(w, fmt.Sprintf("repo: %s\n", repos[i]))
		}
	}()

	if repos != nil {
		return
	}

	user := api.Profile
	repoApi := fmt.Sprintf("https://api.github.com/users/%s/repos?sort=pushed", user)

	resp, err := http.Get(repoApi)
	if err != nil {
		logs.Warn(fmt.Sprintf("error getting repo information: %s\n", err))
	}

	if resp.Body != nil {
		defer resp.Body.Close()
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		logs.Error(err.Error())
	}

	if err = json.Unmarshal(body, &repos); err != nil {
		logs.Error(err.Error())
	}
}
