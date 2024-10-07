package models

import "gorm.io/gorm"

type PointsHistory struct {
	gorm.Model
	RecordID     uint `gorm:"primaryKey"`
	GuildUserID  uint `gorm:"not null"`
	PointsChange int  `gorm:"not null"`

	// Relationships
	GuildUser GuildUser `gorm:"foreignKey:GuildUserID"`
}
