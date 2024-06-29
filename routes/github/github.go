package github

import (
	"fmt"
	"net/http"

	"github.com/a-h/templ"
)

type logger interface {
	Debug(string, ...any)
	Warn(string, ...any)
}

type gitHubApi struct {
	logger  logger
	json    json
	profile string
	repos   []gitHubRepo
}

type gitHubRepo struct {
	RepoName    string `json:"name"`
	Description string `json:"description"`
	Url         string `json:"url"`
}

type json interface {
	UnmarshallUrl(string, any) any
}

func NewRoute(lgr logger, json json, profile string) gitHubApi {
	return gitHubApi{
		logger:  lgr,
		json:    json,
		profile: profile,
	}
}

func (gh gitHubApi) GetRepos() http.Handler {
	// HACK: cache better
	if gh.repos != nil {
		return nil
	}
	const userRepoApi = "https://api.github.com/users/%s/repos?sort=pushed"
	url := fmt.Sprintf(userRepoApi, gh.profile)
	gh.logger.Debug("fetching github user data via", "url", url)
	gh.repos = *gh.json.UnmarshallUrl(url, &gh.repos).(*[]gitHubRepo)
	if len(gh.repos) == 0 {
		gh.logger.Warn("found no repos for", "user", gh.profile, "url", url)
	}
	return templ.Handler(ghProjects(gh.repos))
}
