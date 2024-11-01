package models

import (
	"gorm.io/gorm"
)

type Guild struct {
	gorm.Model
	GuildName string `gorm:"not null"`
	GuildID   string `gorm:"not null"`
	GuildLang string `gorm:"not null" default:"fenix"` // fenix/sage/banshee
}
