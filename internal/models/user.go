package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	UserID   uint   `gorm:"primaryKey"` // Unique identifier for the user
	Username string `gorm:"not null"`   // Name of the user

	// Relationships
	GuildUsers []GuildUser `gorm:"foreignKey:UserID"`
}
