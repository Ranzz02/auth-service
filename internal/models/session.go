package models

import (
	"time"

	gonanoid "github.com/matoous/go-nanoid/v2"
	"gorm.io/gorm"
)

type Session struct {
	Base
	UserID    string    `gorm:"size:21;not null;" json:"user_id"`
	JTI       string    `gorm:"size:21;not null;" json:"token_jti"`
	LastLogin time.Time `gorm:"type:timestamp;" json:"last_login"`
}

func (s *Session) BeforeCreate(tx *gorm.DB) (err error) {
	if s.ID, err = gonanoid.New(); err != nil {
		return err
	}
	return
}
