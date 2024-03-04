package models

import (
	"time"
)

type PhotoModel struct {
	PhotoID   int       `gorm:"primaryKey;autoIncrement" json:"photo_id"`
	UserID    int       `gorm:"foreignKey:UserID"`
	Caption   string    `json:"caption"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
