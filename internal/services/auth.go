package services

import (
	"github.com/Ranzz02/auth-service/internal/models"
	"github.com/Ranzz02/auth-service/internal/utils"
	"github.com/gin-gonic/gin"
)

type AuthService struct {
	repository models.AuthRepository
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
	if _, _, err := a.repository.CreateSession(c, jti); err != nil {
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
