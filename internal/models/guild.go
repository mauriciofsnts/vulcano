package models

import (
	"gorm.io/gorm"
)

type Guild struct {
	gorm.Model
	GuildName string `gorm:"not null"`
	GuildID   string `gorm:"not null"`
}
