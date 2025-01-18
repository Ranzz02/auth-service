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
	Username   string `gorm:"size:20;not null;unique;index" json:"username"`
	Email      string `gorm:"size:50;not null;unique;" json:"email"`
	Password   string `gorm:"not null;" json:"-"`
	Verified   bool   `gorm:"default:false;not null;" json:"-"`
	VerifyCode string `gorm:"not null;" json:"-"`
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

	// Hash users verify code
	if err := u.HashVerifyCode(); err != nil {
		return err
	}

	return
}

// Function hashes the verify code of the user
// Used before saving a new user or when sending a new verify email
func (u *User) HashVerifyCode() error {
	code, err := bcrypt.GenerateFromPassword([]byte(u.VerifyCode), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.VerifyCode = string(code)
	return nil
}

// Verify User checks the code and returns if it's right.
func (u *User) VerifyUser(code string) bool {
	return bcrypt.CompareHashAndPassword([]byte(u.VerifyCode), []byte(code)) == nil
}

// Verify Password checks if the password matches the hash.
func (u *User) VerifyPassword(password string) bool {
	return bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password)) == nil
}

type UserRepository interface {
	GetUser(c *gin.Context, query interface{}, args ...interface{}) (*User, *utils.ApiError, error)
	GetUsers(c *gin.Context, query interface{}, args ...interface{}) (*[]User, *utils.ApiError, error)
	UpdateUser(c *gin.Context, query interface{}, updateData map[string]interface{}, args ...interface{}) (*User, *utils.ApiError, error)
}
