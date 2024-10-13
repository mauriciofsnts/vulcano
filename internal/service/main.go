package service

import "gorm.io/gorm"

type IService struct {
	Guild       IGuildService
	GuildMember IGuildMemberService
	GuildState  IGuildStateService
}

func New(db *gorm.DB) IService {
	guildService := NewGuildService(db)
	guildMemberService := NewGuildMemberService(db)
	guildStateService := NewGuildStateService(db)

	return IService{
		Guild:       guildService,
		GuildMember: guildMemberService,
		GuildState:  guildStateService,
	}
}
