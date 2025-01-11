package repositories

import (
	"github.com/Ranzz02/auth-service/internal/models"
	"github.com/Ranzz02/auth-service/internal/utils"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type AuthRepository struct {
	db *gorm.DB
}

// RegisterUser implements models.AuthRepository.
// This function takes input from a user and creates an entry with that information to the database and returns that user and access & refresh tokens also.
func (a *AuthRepository) RegisterUser(c *gin.Context, registerData *models.RegisterData) (*models.Tokens, *models.User, *utils.ApiError, error) {
	panic("unimplemented")
}

// GetUser implements models.AuthRepository.
func (a *AuthRepository) GetUser(c *gin.Context, query interface{}, args ...interface{}) (*models.Tokens, *models.User, *utils.ApiError, error) {
	panic("unimplemented")
}

// CreateSession implements models.AuthRepository.
func (a *AuthRepository) CreateSession(c *gin.Context) (*models.Tokens, *utils.ApiError, error) {
	panic("unimplemented")
}

// GetSession implements models.AuthRepository.
func (a *AuthRepository) GetSessions(c *gin.Context, query interface{}, args ...interface{}) (*[]models.Session, *utils.ApiError, error) {
	panic("unimplemented")
}

// DeleteSession implements models.AuthRepository.
func (a *AuthRepository) DeleteSession(c *gin.Context) (bool, *utils.ApiError, error) {
	panic("unimplemented")
}

func NewAuthRepository(db *gorm.DB) models.AuthRepository {
	return &AuthRepository{
		db: db,
	}
}
