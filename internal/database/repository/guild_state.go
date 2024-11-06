package repository

import (
	"encoding/json"

	"github.com/mauriciofsnts/bot/internal/database/models"
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

func (r *GuildStateRepository) GetComponentStateByMessageID(id string, entity *models.GuildState) error {
	return r.db.Where("message_id = ?", id).First(entity).Error
}

func (r *GuildStateRepository) UpdateComponentState(messageId string, state map[string]any) error {

	data, err := json.Marshal(state)

	if err != nil {
		return err
	}

	return r.db.Model(&models.GuildState{}).Where("message_id = ?", messageId).Update("state", data).Error
}
