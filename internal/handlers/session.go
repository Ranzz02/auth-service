package handlers

import (
	"net/http"

	"github.com/Ranzz02/auth-service/internal/middleware"
	"github.com/Ranzz02/auth-service/internal/models"
	"github.com/gin-gonic/gin"
)

type SessionHandler struct {
	repository models.AuthRepository
}

func (h *SessionHandler) GetSessions(c *gin.Context) {
	sessions, apiErr, err := h.repository.GetSessions(c, "user_id = ?", middleware.CurrentUser)
	if err != nil {
		c.Error(apiErr)
		return
	}

	c.JSON(http.StatusOK, gin.H{"sessions": sessions})
}

func (h *SessionHandler) DeleteSession(c *gin.Context) {

}

func (h *SessionHandler) DeleteSessions(c *gin.Context) {

}

func NewSessionHandler(router *gin.RouterGroup, repository models.AuthRepository) {
	handler := SessionHandler{
		repository: repository,
	}

	router.GET("", handler.GetSessions)
	router.DELETE("/:id", handler.DeleteSession)
	router.DELETE("", handler.DeleteSessions)
}
