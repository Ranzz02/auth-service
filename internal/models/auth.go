package models

import (
	"github.com/Ranzz02/auth-service/internal/utils"
	"github.com/gin-gonic/gin"
)

type Tokens struct {
	Access  string `json:"access"`
	Refresh string `json:"refresh"`
}

type RegisterData struct {
	Username string `json:"username" validate:"required" binding:"required"`
	Email    string `json:"email" validate:"required" binding:"required"`
	Password string `json:"password" validate:"required" binding:"required"`
}

type AuthRepository interface {
	RegisterUser(c *gin.Context, registerData *RegisterData) (*Tokens, *User, *utils.ApiError, error)
	GetUser(c *gin.Context, query interface{}, args ...interface{}) (*Tokens, *User, *utils.ApiError, error)
	GetSessions(c *gin.Context, query interface{}, args ...interface{}) (*[]Session, *utils.ApiError, error)
	CreateSession(c *gin.Context) (*Tokens, *utils.ApiError, error)
}
