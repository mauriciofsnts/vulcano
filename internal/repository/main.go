package repository

import "gorm.io/gorm"

type IRepository struct {
	Guild GuildRepository
}

func Repositories(db *gorm.DB) IRepository {
	return IRepository{
		Guild: NewGuildRepository(db),
	}
}
