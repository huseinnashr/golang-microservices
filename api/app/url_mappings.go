package app

import (
	"github.com/nvdhunter/golang-microservices/controllers/repository"
	"github.com/nvdhunter/golang-microservices/controllers/repository/polo"
)

func mapUrls() {
	router.GET("/marco", polo.Polo)
	router.POST("/repositories", repository.CreateRepo)
}
