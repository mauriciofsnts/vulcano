package database

import (
	"github.com/mauriciofsnts/bot/internal/models"
	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) error {
	return db.AutoMigrate(&models.User{}, &models.Guild{}, &models.GuildUser{}, &models.PointsHistory{})
}
