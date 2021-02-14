package repository

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nvdhunter/golang-microservices/domain/repository"
	"github.com/nvdhunter/golang-microservices/services"
	"github.com/nvdhunter/golang-microservices/utils/errors"
)

func CreateRepo(c *gin.Context) {
	var request repository.CreateRepoRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		apiErr := errors.NewBadRequestError("invalid json body")
		c.JSON(apiErr.GetStatus(), apiErr)
		return
	}

	result, err := services.RepositoryService.CreateRepo(request)
	if err != nil {
		c.JSON(err.GetStatus(), err)
		return
	}

	c.JSON(http.StatusCreated, result)
}
