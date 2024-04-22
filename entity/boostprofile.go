package entity

import (
	"time"

	"gorm.io/gorm"
)

type BoostProfile struct {
	ID        string         `gorm:"primary_key" json:"id"`
	UserID    string         `gorm:"type:varchar(36);index" json:"user_id"`
	ExpiredAt time.Time      `gorm:"expired_at" json:"expired_at"`
	CreatedAt time.Time      `json:"created_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}
