package service

import "gorm.io/gorm"

type IService struct {
	Guild       IGuildService
	GuildMember IGuildMemberService
}

func New(db *gorm.DB) IService {
	guildService := NewGuildService(db)
	guildMemberService := NewGuildMemberService(db)

	return IService{
		Guild:       guildService,
		GuildMember: guildMemberService,
	}
}
