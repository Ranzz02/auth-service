package repositories

import (
	"github.com/Ranzz02/auth-service/internal/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type AuthRepository struct {
	db *gorm.DB
}

// GetSession implements models.AuthRepository.
func (a *AuthRepository) GetSession(c *gin.Context, query interface{}, args ...interface{}) {
	panic("unimplemented")
}

// GetUser implements models.AuthRepository.
func (a *AuthRepository) GetUser(c *gin.Context, query interface{}, args ...interface{}) (*models.Tokens, *models.User, error) {
	panic("unimplemented")
}

// RegisterUser implements models.AuthRepository.
func (a *AuthRepository) RegisterUser(c *gin.Context, registerData *models.RegisterData) (*models.Tokens, *models.User, error) {
	panic("unimplemented")
}

func NewAuthRepository(db *gorm.DB) models.AuthRepository {
	return &AuthRepository{
		db: db,
	}
}
