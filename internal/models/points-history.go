package models

import "time"

type PointsHistory struct {
	RecordID    uint64    `gorm:"primaryKey"` //TODO! Change this to snowflake
	GuildUserID uint64    `gorm:"not null"`
	CreatedAt   time.Time `gorm:"default:CURRENT_TIMESTAMP"`
	PointsAdded int       `gorm:"not null"`
	TimeSpent   int       `gorm:"not null"`
	GuildUser   GuildUser `gorm:"foreignKey:GuildUserID"`
}
