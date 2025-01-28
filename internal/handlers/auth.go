package handlers

import (
	"fmt"
	"net/http"

	"github.com/Ranzz02/auth-service/internal/middleware"
	"github.com/Ranzz02/auth-service/internal/models"
	"github.com/Ranzz02/auth-service/internal/utils"
	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	repository models.AuthRepository
	service    models.AuthService
}

// Sign In
//
// Function to sign in and create session for user
func (h *AuthHandler) SignIn(c *gin.Context) {
	var input models.SignInData

	// Parse and validate input JSON
	if err := c.ShouldBindBodyWithJSON(&input); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	// Query user by username or email
	user, apiError, err := h.repository.GetUser(c, "username = ? OR email = ?", input.Identity, input.Identity)
	if err != nil {
		c.Error(apiError)
		return
	}

	// Verify password with input
	if !user.VerifyPassword(input.Password) {
		c.Error(utils.InvalidUsernameOrEmail)
		return
	}

	// Generate tokens and saves a session also
	tokens, err := h.service.GenerateTokens(c, user.ID)
	if err != nil {
		c.Error(err)
		return
	}

	// Return successful login with: User and Tokens
	c.JSON(http.StatusOK, gin.H{
		"user":   user,
		"tokens": tokens,
	})
}

// Sign up
//
// Function to try and sign up user and create session
func (h *AuthHandler) SignUp(c *gin.Context) {
	var input models.SignUpData

	// Parse and validate input JSON
	if err := c.ShouldBindBodyWithJSON(&input); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	// Query user by username or email
	user, apiError, err := h.service.SignUpUser(c, input)
	if err != nil {
		c.Error(apiError)
		return
	}

	// Generate tokens and saves a session also
	tokens, err := h.service.GenerateTokens(c, user.ID)
	if err != nil {
		c.Error(err)
		return
	}

	// Return successful login with: User and Tokens
	c.JSON(http.StatusOK, gin.H{
		"user":   user,
		"tokens": tokens,
	})
}

func (h *AuthHandler) SignOut(c *gin.Context) {
	all := c.DefaultQuery("all", "false")
	if all == "true" {
		if ok, apiErr, _ := h.repository.DeleteSessions(c); !ok {
			c.Error(apiErr)
			return
		}
	} else {
		if ok, apiErr, _ := h.repository.DeleteSession(c); !ok {
			c.Error(apiErr)
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{"success": true})
}

func (h *AuthHandler) Refresh(c *gin.Context) {

}

// Verify & Reset
func (h *AuthHandler) ResetPassword(c *gin.Context) {

}

func (h *AuthHandler) ConfirmAccount(c *gin.Context) {
	token := c.Query("token")

	claims := utils.ExtractClaims(token)
	if claims == nil {
		c.Error(fmt.Errorf("Failed to extract claims from token, or token is not valid."))
		return
	}

	user, ok := h.repository.VerifyUser(c, claims["sub"].(string), claims["code"].(string))
	if !ok {
		c.Error(fmt.Errorf("Failed to verify user, try sending a new verify email."))
		return
	}

	c.JSON(http.StatusOK, gin.H{"user": user})
}

func (h *AuthHandler) ResendVerify(c *gin.Context) {

}

// Auth Handler to manage: Signin, Signup, Signout and Refresh
func NewAuthHandler(router *gin.RouterGroup, r models.AuthRepository, s models.AuthService) {
	handler := &AuthHandler{
		repository: r,
		service:    s,
	}

	// Authentication
	router.POST("/signin", handler.SignIn)
	router.POST("/signup", handler.SignUp)
	router.GET("/signout", handler.SignOut)
	router.GET("/refresh", handler.Refresh)

	// Verify & Reset
	router.POST("/reset", handler.ResetPassword)
	router.GET("/resend", middleware.AuthMiddleware, handler.ResendVerify)
	router.GET("/confirm", handler.ConfirmAccount)
}
