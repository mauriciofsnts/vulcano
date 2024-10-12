package repository

import (
	"github.com/mauriciofsnts/bot/internal/models"
	"gorm.io/gorm"
)

type MemberRepository struct {
	*GenericRepository[models.Member]
}

func NewMemberRepository(db *gorm.DB) MemberRepository {
	return MemberRepository{
		NewGenericRepository[models.Member](db),
	}
}

func (r *MemberRepository) GetMemberByUserID(id string, entity *models.Member) error {
	return r.db.Where("user_id = ?", id).First(entity).Error
}
