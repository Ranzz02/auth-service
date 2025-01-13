package models

import "time"

type Session struct {
	Base
	UserID    string    `gorm:"size:21;not null;" json:"user_id"`
	JTI       string    `gorm:"size:21;not null;" json:"token_jti"`
	LastLogin time.Time `gorm:"type:timestamp;" json:"last_login"`
}
