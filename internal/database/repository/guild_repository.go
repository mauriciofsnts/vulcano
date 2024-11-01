package repository

import (
	"github.com/disgoorg/snowflake/v2"
	"github.com/mauriciofsnts/bot/internal/database/models"
	"gorm.io/gorm"
)

type GuildRepository struct {
	*GenericRepository[models.Guild]
}

func NewGuildRepository(db *gorm.DB) GuildRepository {
	return GuildRepository{
		NewGenericRepository[models.Guild](db),
	}
}

func (r *GuildRepository) GetGuildByGuildId(id snowflake.ID, entity *models.Guild) error {
	return r.db.Where("guild_id = ?", id).First(entity).Error
}
