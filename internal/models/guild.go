package models

import "time"

type Guild struct {
	GuildID   uint      `gorm:"primaryKey"`
	GuildName string    `gorm:"not null"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP"`

	// Relationships
	GuildUsers []GuildUser `gorm:"foreignKey:GuildID"`
}
