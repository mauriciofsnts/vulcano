package service

import (
	"github.com/disgoorg/snowflake/v2"
	"github.com/mauriciofsnts/bot/internal/repository"
	"gorm.io/gorm"
)

type IGuildService interface {
	EnsureGuildExists(guildId snowflake.ID)
}

type GuildService struct {
	repository repository.GuildRepository
}

func NewGuildService(db *gorm.DB) IGuildService {
	return &GuildService{
		repository: repository.NewGuildRepository(db),
	}
}

func (s *GuildService) EnsureGuildExists(guildId snowflake.ID) {
}
