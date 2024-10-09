package repository

import (
	"github.com/mauriciofsnts/bot/internal/models"
	"gorm.io/gorm"
)

type GuildRepository struct {
	*GenericRepository[models.Guild]
}

func NewGuildRepository(db *gorm.DB) GuildRepository {
	return GuildRepository{
		GenericRepository: NewGenericRepository[models.Guild](db),
	}
}
