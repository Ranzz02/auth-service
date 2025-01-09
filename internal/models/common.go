package models

import (
	"time"

	"gorm.io/gorm"
)

type Base struct {
	ID        string         `gorm:"primaryKey;size:21;" json:"id"` // Custom ID field using nanoid
	CreatedAt time.Time      `gorm:"not null;default:'now()';" json:"createdAt"`
	UpdatedAt time.Time      `gorm:"not null;default:'now()';" json:"updatedAt"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}
