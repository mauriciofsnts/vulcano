package service

import "gorm.io/gorm"

type IService struct {
	Guild IGuildService
}

func New(db *gorm.DB) IService {
	guildService := NewGuildService(db)

	return IService{
		Guild: guildService,
	}
}
