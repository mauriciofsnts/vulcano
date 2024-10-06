package models

import "time"

type User struct {
	UserID    uint64    `gorm:"primaryKey"` //TODO! Change this to snowflake
	Username  string    `gorm:"not null"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP"`
}
