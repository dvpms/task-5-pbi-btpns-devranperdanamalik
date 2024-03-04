package models

import (
	"time"
)

type UserModel struct {
	UserID    int          `gorm:"primaryKey;autoIncrement" json:"user_id"`
	Username  string       `gorm:"type:varchar(200)" binding:"required" json:"username"`
	Email     string       `gorm:"type:varchar(200)" json:"email"`
	Password  string       `gorm:"type:varchar(100)" binding:"required,min=6" json:"password"`
	Photos    []PhotoModel `json:"photos" gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	CreatedAt time.Time    `gorm:"type:datetime" json:"created_at"`
	UpdatedAt time.Time    `gorm:"type:datetime" json:"updated_at"`
}
