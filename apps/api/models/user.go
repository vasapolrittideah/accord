package models

import (
	"github.com/google/uuid"
	"time"
)

type User struct {
	ID           uuid.UUID `json:"id" gorm:"primaryKey;type:uuid;default:uuid_generate_v4()"`
	Name         string    `json:"name" gorm:"type:varchar(128)"`
	Email        string    `json:"email" gorm:"uniqueIndex;not null"`
	Password     string    `json:"-" gorm:"not null"`
	Role         string    `json:"role" gorm:"type:varchar(16);not null"`
	Provider     string    `json:"provider" gorm:"type:varchar(128);not null"`
	Avatar       string    `json:"avatar,omitempty"`
	Verified     bool      `json:"verified"`
	CreatedAt    time.Time `json:"created_at" gorm:"autoCreateTime;not null"`
	UpdatedAt    time.Time `json:"updated_at" gorm:"autoUpdateTime:milli;not null"`
	RefreshToken string    `json:"-"`
}
