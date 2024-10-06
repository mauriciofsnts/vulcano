package models

import "time"

type User struct {
	UserID    uint      `gorm:"primaryKey"`                // Unique identifier for the user
	Username  string    `gorm:"not null"`                  // Name of the user
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP"` // Date when the user was created

	// Relationships
	GuildUsers []GuildUser `gorm:"foreignKey:UserID"`
}
