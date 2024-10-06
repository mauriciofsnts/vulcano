package models

type GuildUser struct {
	GuildUserID uint64 `gorm:"primaryKey"` //TODO! Change this to snowflake
	GuildID     uint64 `gorm:"not null"`
	UserID      uint64 `gorm:"not null"`
	Points      int    `gorm:"default:0"`

	User  User  `gorm:"foreignKey:UserID"`
	Guild Guild `gorm:"foreignKey:GuildID"`
}
