package handlers

import (
	"net/http"

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
func (h *AuthHandler) Signin(c *gin.Context) {
	var input models.SignInData

	// Parse and validate input JSON
	if err := c.ShouldBindBodyWithJSON(&input); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	// Query user by username or email
	user, apiError, err := h.repository.GetUser(c, "WHERE username = ? OR email = ?", input.Identity, input.Identity)
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
func (h *AuthHandler) Signup(c *gin.Context) {
	var input models.SignUpData

	// Parse and validate input JSON
	if err := c.ShouldBindBodyWithJSON(&input); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	// Query user by username or email
	user, apiError, err := h.repository.CreateUser(c, &input)
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

func (h *AuthHandler) Signout(c *gin.Context) {

}

func (h *AuthHandler) Refresh(c *gin.Context) {

}

// Verify & Reset
func (h *AuthHandler) ResetPassword(c *gin.Context) {

}

func (h *AuthHandler) Verify(c *gin.Context) {

}

// Auth Handler to manage: Signin, Signup, Signout and Refresh
func NewAuthHandler(router *gin.RouterGroup, r models.AuthRepository, s models.AuthService) {
	handler := &AuthHandler{
		repository: r,
		service:    s,
	}

	// Authentication
	router.POST("/signin", handler.Signin)
	router.POST("/signup", handler.Signup)
	router.GET("/signout", handler.Signout)
	router.GET("/refresh", handler.Refresh)

	// Verify & Reset
	router.POST("/reset", handler.ResetPassword)
	router.GET("/verify", handler.Verify)
}
