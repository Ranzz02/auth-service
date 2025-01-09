package models

type Session struct {
	ID     string `gorm:"primaryKey;size:21;" json:"id"`
	UserID string `gorm:"size:21;not null;" json:"user_id"`
	JTI    string `gorm:"size:21;not null;" json:"token_jti"`
}
