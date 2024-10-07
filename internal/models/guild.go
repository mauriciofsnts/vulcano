package models

import (
	"time"

	"gorm.io/gorm"
)

type Guild struct {
	gorm.Model
	GuildName string    `gorm:"not null"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP"`

	// Relationships
	GuildUsers []GuildUser `gorm:"foreignKey:GuildID"`
}
