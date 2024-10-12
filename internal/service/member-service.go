package service

import (
	"github.com/disgoorg/disgo/discord"
	"github.com/mauriciofsnts/bot/internal/models"
	"github.com/mauriciofsnts/bot/internal/repository"
	"gorm.io/gorm"
)

type IMemberService interface{}

type MemberService struct {
	repository repository.MemberRepository
}

func NewMemberService(db *gorm.DB) IMemberService {
	return &MemberService{
		repository: repository.NewMemberRepository(db),
	}
}

func (s *MemberService) EnsureMemberExists(member discord.Member) bool {
	m := &models.Member{}
	err := s.repository.GetMemberByUserID(member.User.ID.String(), m)

	if err != nil {
		err := s.repository.Create(&models.Member{
			UserID:   member.User.ID.String(),
			Username: member.User.Username,
		})

		return err == nil
	}

	return true
}
