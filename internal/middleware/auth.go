package middleware

import (
	"github.com/Ranzz02/auth-service/internal/utils"
	"github.com/gin-gonic/gin"
)

var CurrentUser *string

func AuthMiddleware(c *gin.Context) {
	CurrentUser = nil

	token := utils.ExtractToken(c)
	if token == nil {
		c.Error(utils.NoTokenProvided)
		return
	}

	claims := utils.ExtractClaims(*token)
	if claims == nil {
		c.Error(utils.InvalidTokenProvided)
		return
	}

	sub, ok := claims["sub"].(string)
	if !ok {
		c.Error(utils.InvalidTokenProvided)
		return
	}

	CurrentUser = &sub

	c.Next()
}
