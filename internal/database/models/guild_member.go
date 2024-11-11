package models

import "gorm.io/gorm"

type GuildMember struct {
	gorm.Model
	GuildID      string `gorm:"not null"`
	MemberID     string `gorm:"not null"`
	MessageCount uint   `gorm:"default:0"`
	CommandCount uint   `gorm:"default:0"`
	Coins        uint   `gorm:"default:0"`
}
