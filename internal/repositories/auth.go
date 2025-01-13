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
//
// This function takes input from a user and creates an entry with that information to the database and returns that user and access & refresh tokens also.
func (a *AuthRepository) CreateUser(c *gin.Context, registerData *models.SignUpData) (*models.User, *utils.ApiError, error) {
	// start transaction
	tx := a.db.Begin()

	user := &models.User{
		Username: registerData.Username,
		Email:    registerData.Email,
		Password: registerData.Password,
	}

	// Try to create instance in database
	res := tx.Model(&models.User{}).Create(user)
	if res.Error != nil {
		tx.Rollback() // Rollback database
		return nil, &utils.UsernameOrEmailInUse, res.Error
	}

	// Commit changes to the database
	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return nil, &utils.InternalServerError, err
	}

	return user, nil, nil
}

// GetUser implements models.AuthRepository.
//
// This function is used to login a user if found.
func (a *AuthRepository) GetUser(c *gin.Context, query interface{}, args ...interface{}) (*models.User, *utils.ApiError, error) {
	var user models.User

	// Try to locate the user
	if err := a.db.Model(&models.User{}).Where(query, args).First(&user).Error; err != nil {
		return nil, &utils.ResourceNotFound, err
	}

	return &user, nil, nil
}

// CreateSession implements models.AuthRepository.
func (a *AuthRepository) CreateSession(c *gin.Context, jti string) (*models.Session, *utils.ApiError, error) {
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
