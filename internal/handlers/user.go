package handlers

import (
	"github.com/Ranzz02/auth-service/internal/models"
	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	repository *models.UserRepository
}

func (h *UserHandler) GetUsers(c *gin.Context) {

}

func (h *UserHandler) GetUser(c *gin.Context) {

}

func NewUserHandler(router *gin.RouterGroup, repository models.UserRepository) {
	handler := UserHandler{
		repository: &repository,
	}

	router.GET("/users", handler.GetUsers)    // Get users
	router.GET("/users/:id", handler.GetUser) // Get user
}
