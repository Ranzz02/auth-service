package handlers

import (
	"github.com/Ranzz02/auth-service/internal/models"
	"github.com/gin-gonic/gin"
)

type SessionHandler struct {
	repository *models.AuthRepository
}

func (h *SessionHandler) GetSessions(c *gin.Context) {

}

func (h *SessionHandler) DeleteSession(c *gin.Context) {

}

func NewSessionHandler(router *gin.RouterGroup, repository models.AuthRepository) {
	handler := SessionHandler{
		repository: &repository,
	}

	router.GET("/sessions", handler.GetSessions)
	router.DELETE("/sessions/:id", handler.DeleteSession)
}
