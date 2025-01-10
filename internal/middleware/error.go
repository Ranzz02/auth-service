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
		lastErr := (c.Errors.Last().Err).(utils.ApiError)
		for _, err := range c.Errors {
			if apiErr, ok := err.Err.(utils.ApiError); ok {
				apiErrors = append(apiErrors, apiErr)
			}
		}

		if len(apiErrors) > 0 {
			c.AbortWithStatusJSON(lastErr.StatusCode, gin.H{"errors": apiErrors})
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
