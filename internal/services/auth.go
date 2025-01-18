package services

import (
	"fmt"

	"github.com/Ranzz02/auth-service/internal/models"
	"github.com/Ranzz02/auth-service/internal/utils"
	"github.com/gin-gonic/gin"
)

type AuthService struct {
	repository models.AuthRepository
}

func (a *AuthService) SignUpUser(c *gin.Context, signUpData models.SignUpData) (*models.User, *utils.ApiError, error) {
	tx := a.repository.GetDB().Begin()

	user, verifyCode, apiError, err := a.repository.CreateUser(c, signUpData, tx)
	if err != nil {
		tx.Rollback()
		return nil, apiError, err
	}

	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return nil, &utils.InternalServerError, err
	}

	token, err := utils.GenerateVerifyToken(user.ID, verifyCode)
	if err != nil {
		return nil, &utils.InternalServerError, err
	}
	go utils.SendConfirmEmail(utils.ConfirmMailOptions{Username: user.Username, Link: fmt.Sprintf("http://localhost:3000/auth/confirm?token=%s", token), To: user.Email})

	return user, nil, err
}

// GenerateTokens implements models.AuthService.
func (a *AuthService) GenerateTokens(c *gin.Context, id string) (*models.Tokens, error) {
	// Access token
	access, err := utils.GenerateAccessToken(id)
	if err != nil {
		return nil, err
	}

	// Refresh token
	jti, refresh, err := utils.GenerateRefreshToken(id)
	if err != nil {
		return nil, err
	}

	// Save session to database
	if _, _, err := a.repository.CreateSession(c, id, jti); err != nil {
		return nil, err
	}

	return &models.Tokens{
		Access:  access,
		Refresh: refresh,
	}, nil
}

func NewAuthService(repository models.AuthRepository) models.AuthService {
	return &AuthService{
		repository: repository,
	}
}
