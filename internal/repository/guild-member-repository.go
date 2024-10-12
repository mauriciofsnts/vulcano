package repository

import (
	"github.com/mauriciofsnts/bot/internal/models"
	"gorm.io/gorm"
)

type GuildMemberRepository struct {
	*GenericRepository[models.GuildMember]
}

func NewGuildMemberRepository(db *gorm.DB) GuildMemberRepository {
	return GuildMemberRepository{
		NewGenericRepository[models.GuildMember](db),
	}
}

func (r *GuildMemberRepository) GetGuildMemberByUserID(id string, entity *models.GuildMember) error {
	return r.db.Where("member_id = ?", id).First(entity).Error
}

func (r *GuildMemberRepository) GetGuildMemberByGuildIDAndUserID(
	guildID, userID string,
	entity *models.GuildMember,
) error {
	return r.db.Where("guild_id = ? AND member_id = ?", guildID, userID).First(entity).Error
}
