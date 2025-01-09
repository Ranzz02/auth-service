package handlers

import (
	"github.com/Ranzz02/auth-service/internal/models"
	"github.com/Ranzz02/auth-service/internal/utils"
	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	repository *models.AuthRepository
}

func (h *AuthHandler) Signin(c *gin.Context) {
	c.Error(utils.InvalidUsernameOrEmail)
}

func (h *AuthHandler) Signup(c *gin.Context) {

}

func (h *AuthHandler) Signout(c *gin.Context) {

}

func (h *AuthHandler) Refresh(c *gin.Context) {

}

// Auth Handler to manage: Signin, Signup, Signout and Refresh
func NewAuthHandler(router *gin.RouterGroup, r models.AuthRepository) {
	handler := &AuthHandler{
		repository: &r,
	}

	router.POST("/signin", handler.Signin)
	router.POST("/signup", handler.Signup)
	router.GET("/signout", handler.Signout)
}
