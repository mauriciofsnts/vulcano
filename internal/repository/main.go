package repository

import "gorm.io/gorm"

type IRepository struct {
	Guild       GuildRepository
	Member      MemberRepository
	GuildMember GuildMemberRepository
}

func Repositories(db *gorm.DB) IRepository {
	return IRepository{
		Guild:       NewGuildRepository(db),
		Member:      NewMemberRepository(db),
		GuildMember: NewGuildMemberRepository(db),
	}
}
