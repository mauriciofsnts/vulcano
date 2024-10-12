package service

import "gorm.io/gorm"

type IService struct {
	Guild       IGuildService
	Member      IMemberService
	GuildMember IGuildMemberService
}

func New(db *gorm.DB) IService {
	guildService := NewGuildService(db)
	memberService := NewMemberService(db)
	guildMemberService := NewGuildMemberService(db)

	return IService{
		Guild:       guildService,
		Member:      memberService,
		GuildMember: guildMemberService,
	}
}
