package models

import (
	"github.com/Ranzz02/auth-service/internal/utils"
	"github.com/gin-gonic/gin"
)

type Tokens struct {
	Access  string `json:"access"`
	Refresh string `json:"refresh"`
}

type SignInData struct {
	Identity string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type SignUpData struct {
	Username string `json:"username" validate:"required" binding:"required"`
	Email    string `json:"email" validate:"required,email" binding:"required,email"`
	Password string `json:"password" validate:"required" binding:"required"`
}

type AuthRepository interface {
	CreateUser(c *gin.Context, registerData *SignUpData) (*User, *utils.ApiError, error)
	GetUser(c *gin.Context, query interface{}, args ...interface{}) (*User, *utils.ApiError, error)
	GetSessions(c *gin.Context, query interface{}, args ...interface{}) (*[]Session, *utils.ApiError, error)
	CreateSession(c *gin.Context, jti string) (*Session, *utils.ApiError, error)
	DeleteSession(c *gin.Context) (bool, *utils.ApiError, error)
}

type AuthService interface {
	GenerateTokens(c *gin.Context, id string) (*Tokens, error)
}
