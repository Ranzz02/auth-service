package middleware

import (
	"net/http"

	"github.com/Ranzz02/auth-service/internal/utils"
	"github.com/gin-gonic/gin"
)

func ErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		var apiErrors []utils.ApiError
		for _, err := range c.Errors {
			if apiErr, ok := err.Err.(utils.ApiError); ok {
				apiErrors = append(apiErrors, apiErr)
			}
		}

		if len(apiErrors) > 0 {
			statusCode := http.StatusBadRequest
			switch apiErrors[len(apiErrors)-1].Code {
			case 101, 102:
				statusCode = http.StatusUnauthorized
			case 301:
				statusCode = http.StatusNotFound
			case 302:
				statusCode = http.StatusConflict
			}

			c.AbortWithStatusJSON(statusCode, gin.H{"errors": apiErrors})
			return
		}

		if len(c.Errors) > 0 {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"code":    500,
				"message": "Internal servver error",
			})
			return
		}
	}
}