package models

import "github.com/gin-gonic/gin"

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
	RegisterUser(c *gin.Context, registerData *RegisterData) (*Tokens, *User, error)
	GetUser(c *gin.Context, query interface{}, args ...interface{}) (*Tokens, *User, error)
	GetSession(c *gin.Context, query interface{}, args ...interface{}) 
}
