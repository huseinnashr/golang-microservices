package config

import "os"

const (
	apiGithubAccessToken = "SECRET_GITHUB_ACESS_TOKEN"
)

var (
	githubAccessToken = os.Getenv(apiGithubAccessToken)
)

func GetGithubAccessToken() string {
	return githubAccessToken
}
