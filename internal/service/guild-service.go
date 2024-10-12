package service

import (
	"github.com/disgoorg/disgo/discord"
	"github.com/mauriciofsnts/bot/internal/models"
	"github.com/mauriciofsnts/bot/internal/repository"
	"gorm.io/gorm"
)

type IGuildService interface {
	EnsureGuildExists(guild discord.Guild) bool
}

type GuildService struct {
	repository repository.GuildRepository
}

func NewGuildService(db *gorm.DB) IGuildService {
	return &GuildService{
		repository: repository.NewGuildRepository(db),
	}
}

func (s *GuildService) EnsureGuildExists(guild discord.Guild) bool {
	g := &models.Guild{}
	err := s.repository.GetGuildByGuildId(guild.ID, g)

	if err != nil {
		err := s.repository.Create(&models.Guild{
			GuildName: guild.Name,
			GuildID:   guild.ID.String(),
		})

		return err == nil
	}

	return true
}
