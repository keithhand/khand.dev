package handlers

import (
	"fmt"
	"log/slog"
)

type Json interface {
	UnmarshallUrl(string, any) any
}

type Config interface {
	MockApi() bool
	GhProfile() string
}

type gitHubRepo struct {
	RepoName    string `json:"name"`
	Description string `json:"description"`
	RepoUrl     string `json:"url"`
}

type gitHubApi struct {
	logger  Logger
	json    Json
	profile string
	repos   []gitHubRepo
}

func GitHub(config Config, json Json) *handler {
	gh := &gitHubApi{
		logger:  slog.Default(),
		json:    json,
		profile: config.GhProfile(),
	}
	if config.MockApi() {
		gh.repos = mockRepos()
	}
	return New(gh)
}

func (gh *gitHubApi) Repos() []gitHubRepo {
	// HACK: cache better
	if len(gh.repos) != 0 {
		gh.logger.Warn("cached repos found")
		return gh.repos
	}
	url := gh.getUserReposUrl()
	gh.logger.Debug("fetching github user data via", "url", url)
	gh.json.UnmarshallUrl(url, &gh.repos)
	if len(gh.repos) == 0 {
		gh.logger.Warn("found no repos for", "user", gh.profile, "url", url)
	}
	return gh.repos
}

func (gh gitHubApi) getUserReposUrl() string {
	const userRepoApi = "https://api.github.com/users/%s/repos?sort=pushed"
	return fmt.Sprintf(userRepoApi, gh.profile)
}

func mockRepos() []gitHubRepo {
	var ghr []gitHubRepo
	var mockRepos []string = []string{"khand.dev", "dotfiles", "resume", "nvim",
		"app-msngr", "keithhand", "homelab", "runelite-time-tracking-reminder"}
	for _, repo := range mockRepos {
		ghr = append(ghr, gitHubRepo{
			RepoName:    repo,
			Description: fmt.Sprintf("%s repo, but mocked.", repo),
			RepoUrl:     fmt.Sprintf("https://api.github.com/repos/keithhand/%s", repo),
		})
	}
	return ghr
}
