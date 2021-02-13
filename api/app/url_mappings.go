package app

import (
	"github.com/nvdhunter/golang-microservices/controllers/repository"
)

func mapUrls() {
	router.POST("/repositories", repository.CreateRepo)
}
