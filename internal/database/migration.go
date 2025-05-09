package database

import (
	"github.com/mauriciofsnts/bot/internal/database/models"
	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) error {
	return db.AutoMigrate(&models.Guild{}, &models.GuildMember{})
}
