package repository

import "gorm.io/gorm"

type IRepository struct {
	Guild       GuildRepository
	GuildMember GuildMemberRepository
	GuildState  GuildStateRepository
}

func Repositories(db *gorm.DB) IRepository {
	return IRepository{
		Guild:       NewGuildRepository(db),
		GuildMember: NewGuildMemberRepository(db),
		GuildState:  NewGuildStateRepository(db),
	}
}
