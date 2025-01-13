package handlers

import (
	"github.com/Ranzz02/auth-service/internal/models"
	"github.com/gin-gonic/gin"
)

type OAuthHandler struct {
	repository models.OAuthRepository
}

func (h *OAuthHandler) RedirectToProvider(c *gin.Context) {

}

func (h *OAuthHandler) ProviderCallback(c *gin.Context) {

}

func NewOAuthHandler(router *gin.RouterGroup, repository models.OAuthRepository) {
	handler := &OAuthHandler{
		repository: repository,
	}

	router.GET("/auth/:provider", handler.RedirectToProvider)
	router.GET("/auth/:provider/callback", handler.ProviderCallback)
}
