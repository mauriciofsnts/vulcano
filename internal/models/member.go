package models

import "gorm.io/gorm"

type Member struct {
	gorm.Model
	UserID   string `gorm:"primaryKey"` // Unique identifier for the user
	Username string `gorm:"not null"`   // Name of the user
}
