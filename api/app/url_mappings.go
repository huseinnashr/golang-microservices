package app

import (
	"github.com/nvdhunter/golang-microservices/controllers/polo"
	"github.com/nvdhunter/golang-microservices/controllers/repository"
)

func mapUrls() {
	router.GET("/marco", polo.Marco)
	router.POST("/repositories", repository.CreateRepo)
}
