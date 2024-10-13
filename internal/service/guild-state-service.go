package service

import (
	"github.com/mauriciofsnts/bot/internal/models"
	"github.com/mauriciofsnts/bot/internal/repository"
	"gorm.io/gorm"
)

type IGuildStateService interface {
	GetComponentStateById(id string) (*models.GuildState, error)
	UpdateComponentState(id string, state []any) error
	CreateComponentState(guildState *models.GuildState) error
}

type GuildStateService struct {
	repository repository.GuildStateRepository
}

func NewGuildStateService(db *gorm.DB) IGuildStateService {
	return &GuildStateService{
		repository: repository.NewGuildStateRepository(db),
	}
}

func (r *GuildStateService) GetComponentStateById(id string) (*models.GuildState, error) {
	guildState := &models.GuildState{}

	err := r.repository.GetComponentStateByID(id, guildState)
	return guildState, err
}

func (r *GuildStateService) UpdateComponentState(id string, state []any) error {
	return r.repository.UpdateComponentState(id, state)
}

func (r *GuildStateService) CreateComponentState(guildState *models.GuildState) error {
	return r.repository.Create(guildState)
}
