package models

import "time"

type Guild struct {
	GuildID   uint64    `gorm:"primaryKey"` //TODO! Change this to snowflake
	GuildName string    `gorm:"not null"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP"`
}
