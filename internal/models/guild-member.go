package models

import "gorm.io/gorm"

type GuildMember struct {
	gorm.Model
	GuildID      string `gorm:"primaryKey"` // Unique identifier for the guild
	MemberID     string `gorm:"primaryKey"` // Unique identifier for the user
	MessageCount uint   `gorm:"default:0"`  // Number of messages sent by the user in the guild
	CommandCount uint   `gorm:"default:0"`  // Number of commands executed by the user in the guild
	Coins        uint   `gorm:"default:0"`  // Number of coins the user has in the guild

	// Relationships
	Guild Guild `gorm:"foreignKey:GuildID"`
}
