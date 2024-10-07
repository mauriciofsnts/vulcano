package models

import "gorm.io/gorm"

type GuildUser struct {
	gorm.Model
	UserID  uint `gorm:"not null"`
	GuildID uint `gorm:"not null"`

	// Relationships
	User          User            `gorm:"foreignKey:UserID"`
	Guild         Guild           `gorm:"foreignKey:GuildID"`
	PointsHistory []PointsHistory `gorm:"foreignKey:GuildUserID"`
}
