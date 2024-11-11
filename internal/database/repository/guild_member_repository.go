package repository

import (
	"github.com/mauriciofsnts/bot/internal/database/models"
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
	return r.db.Where("member_id = ?", id).Find(entity).Error
}

func (r *GuildMemberRepository) GetGuildMemberByGuildIDAndUserID(
	guildID, userID string,
	entity *models.GuildMember,
) error {
	return r.db.Where(&models.GuildMember{GuildID: guildID, MemberID: userID}).First(entity).Error
}
