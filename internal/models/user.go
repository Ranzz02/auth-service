package models

import (
	"github.com/Ranzz02/auth-service/internal/utils"
	"github.com/gin-gonic/gin"
	gonanoid "github.com/matoous/go-nanoid/v2"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	Base
	Username string `gorm:"size:20;not null;unique;index" json:"username"`
	Email    string `gorm:"size:50;not null;unique;" json:"email"`
	Password string `gorm:"not null;" json:"-"`
	Verified bool   `gorm:"default:false;not null;" json:"-"`
	Profile
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	if u.ID, err = gonanoid.New(); err != nil {
		return err
	}
	// Hash users password
	hash, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.Password = string(hash)
	return
}

func (u *User) VerifyPassword(password string) bool {
	return bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password)) == nil
}

type UserRepository interface {
	GetUsers(c *gin.Context, query interface{}, args ...interface{}) (*User, *utils.ApiError, error)
	UpdateUser(c *gin.Context, query interface{}, updateData map[string]interface{}, args ...interface{}) (*User, *utils.ApiError, error)
}
