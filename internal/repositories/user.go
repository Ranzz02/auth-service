package repositories

import (
	"github.com/Ranzz02/auth-service/internal/models"
	"github.com/Ranzz02/auth-service/internal/utils"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

// GetUser implements models.UserRepository.
func (r *UserRepository) GetUser(c *gin.Context, query interface{}, args ...interface{}) (*models.User, *utils.ApiError, error) {
	panic("unimplemented")
}

// UpdateUser implements models.UserRepository.
func (r *UserRepository) UpdateUser(c *gin.Context, query interface{}, updateData map[string]interface{}, args ...interface{}) (*models.User, *utils.ApiError, error) {
	panic("unimplemented")
}

// GetUsers implements models.UserRepository.
func (r *UserRepository) GetUsers(c *gin.Context, query interface{}, args ...interface{}) (*[]models.User, *utils.ApiError, error) {
	panic("unimplemented")
}

func NewUserRepository(db *gorm.DB) models.UserRepository {
	return &UserRepository{
		db: db,
	}
}
