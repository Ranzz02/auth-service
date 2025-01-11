package handlers

import (
	"github.com/Ranzz02/auth-service/internal/models"
	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	repository *models.UserRespository
}

func (h *UserHandler) GetUsers(c *gin.Context) {

}

func (h *UserHandler) GetUser(c *gin.Context) {

}

func NewUserHandler(router *gin.RouterGroup, repository models.UserRespository) {
	handler := UserHandler{
		repository: &repository,
	}

	router.GET("/users", handler.GetUsers)    // Get users
	router.GET("/users/:id", handler.GetUser) // Get user
}
