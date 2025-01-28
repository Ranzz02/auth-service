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
	var user models.User

	if err := r.db.Model(&models.User{}).Where(query, args).First(&user).Error; err != nil {
		return nil, &utils.ResourceNotFound, err
	}

	return &user, nil, nil
}

// UpdateUser implements models.UserRepository.
func (r *UserRepository) UpdateUser(c *gin.Context, query interface{}, updateData map[string]interface{}, args ...interface{}) (*models.User, *utils.ApiError, error) {
	var user models.User

	if err := r.db.Model(&models.User{}).Where(query, args).Updates(updateData).First(&user).Error; err != nil {
		return nil, &utils.ResourceConflict, err
	}

	return &user, nil, nil
}

// GetUsers implements models.UserRepository.
func (r *UserRepository) GetUsers(c *gin.Context, query interface{}, args ...interface{}) (*[]models.User, *utils.ApiError, error) {
	var users []models.User

	if err := r.db.Model(&models.User{}).Where(query, args).Find(&users).Error; err != nil {
		return nil, &utils.ResourceConflict, err
	}

	return &users, nil, nil
}

func NewUserRepository(db *gorm.DB) models.UserRepository {
	return &UserRepository{
		db: db,
	}
}
