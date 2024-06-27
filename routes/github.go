package routes

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

const userRepoApi = "https://api.github.com/users/%s/repos?sort=pushed"

type gitHubApi struct {
	logger  logger
	profile string
	repos   []gitHubRepo
}

type gitHubRepo struct {
	RepoName    string `json:"name"`
	Description string `json:"description"`
	Url         string `json:"url"`
}

func NewGitHub(lgr logger, profile string) gitHubApi {
	return gitHubApi{
		logger:  lgr,
		profile: profile,
	}
}

func (gh gitHubApi) GetRepos(w http.ResponseWriter, r *http.Request) {
	defer func() {
		for i := range gh.repos {
			io.WriteString(w, fmt.Sprintf("repo: %s\n", gh.repos[i]))
		}
	}()

	if gh.repos != nil {
		return
	}

	getUserRepos(&gh.repos, gh.profile, gh.logger)
}

func getUserRepos(rr *[]gitHubRepo, user string, log logger) {
	url := fmt.Sprintf(userRepoApi, user)
	log.Debug("fetching github user data via", "url", url)

	resp, err := http.Get(url)
	if err != nil {
		log.Error(fmt.Errorf("initializing request to %s: %w", url, err).Error())
	}

	if resp.Body != nil {
		defer resp.Body.Close()
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Error(fmt.Errorf("reading repo reponse: %w", err).Error())
	}

	if err = json.Unmarshal(body, rr); err != nil {
		log.Error(fmt.Errorf("unmarshalling repo json: %w", err).Error())
	}

	if len(*rr) == 0 {
		log.Warn("found no repos for", "user", user)
	}
}
