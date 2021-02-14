package services

import (
	"strings"

	"github.com/nvdhunter/golang-microservices/config"
	"github.com/nvdhunter/golang-microservices/domain/github"
	"github.com/nvdhunter/golang-microservices/domain/repository"
	"github.com/nvdhunter/golang-microservices/providers/github_provider"
	"github.com/nvdhunter/golang-microservices/utils/errors"
)

type repoService struct{}

type repoServiceInterface interface {
	CreateRepo(request repository.CreateRepoRequest) (repository.CreateRepoResponse, errors.ApiError)
}

var (
	RepositoryService *repoService
)

func init() {
	RepositoryService = &repoService{}
}

func (s *repoService) CreateRepo(input repository.CreateRepoRequest) (*repository.CreateRepoResponse, errors.ApiError) {
	input.Name = strings.TrimSpace(input.Name)
	if input.Name == "" {
		return nil, errors.NewBadRequestError("invalid repository name")
	}

	request := github.CreateRepoRequest{
		Name:        input.Name,
		Description: input.Description,
		Private:     false,
	}

	response, err := github_provider.CreateRepo(config.GetGithubAccessToken(), request)
	if err != nil {
		return nil, errors.NewApiError(err.StatusCode, err.Message)
	}

	result := repository.CreateRepoResponse{
		Id:    response.Id,
		Name:  response.Name,
		Owner: response.Owner.Login,
	}
	return &result, nil
}
