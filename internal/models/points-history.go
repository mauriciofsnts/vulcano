package models

import "time"

type PointsHistory struct {
	RecordID     uint      `gorm:"primaryKey"`
	GuildUserID  uint      `gorm:"not null"`
	Date         time.Time `gorm:"default:CURRENT_TIMESTAMP"`
	PointsChange int       `gorm:"not null"`

	// Relationships
	GuildUser GuildUser `gorm:"foreignKey:GuildUserID"`
}
