package repository

import (
	"github.com/mauriciofsnts/bot/internal/models"
	"gorm.io/gorm"
)

type GuildStateRepository struct {
	*GenericRepository[models.GuildState]
}

func NewGuildStateRepository(db *gorm.DB) GuildStateRepository {
	return GuildStateRepository{
		NewGenericRepository[models.GuildState](db),
	}
}

func (r *GuildStateRepository) GetComponentStateByID(id string, entity *models.GuildState) error {
	return r.db.Where("component_id = ?", id).First(entity).Error
}

func (r *GuildStateRepository) UpdateComponentState(id string, state []interface{}) error {
	return r.db.Model(&models.GuildState{}).Where("component_id = ?", id).Update("state", state).Error
}