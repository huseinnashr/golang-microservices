package app

import (
	"github.com/nvdhunter/golang-microservices/controllers/polo"
	"github.com/nvdhunter/golang-microservices/controllers/repository"
)

func mapUrls() {
	router.GET("/marco", polo.Polo)
	router.POST("/repositories", repository.CreateRepo)
}
