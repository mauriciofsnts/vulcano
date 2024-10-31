package service

import (
	"log/slog"

	"github.com/mauriciofsnts/bot/internal/models"
	"github.com/mauriciofsnts/bot/internal/repository"
	"gorm.io/gorm"
)

type IGuildMemberService interface {
	EnsureMemberValidity(guildID, userID string) error
	IncrementMessageCount(guildID, userID string)
	IncrementCommandCount(guildID, userID string)
	GetBalance(guildID, userID string) (uint, error)
}

type GuildMemberService struct {
	repository repository.GuildMemberRepository
}

func NewGuildMemberService(db *gorm.DB) IGuildMemberService {
	return &GuildMemberService{
		repository: repository.NewGuildMemberRepository(db),
	}
}

// Verify if the user is a member of the guild if not, check if the users has already created on member table
// if not, create a new member
func (r *GuildMemberService) EnsureMemberValidity(guildID, userID string) error {
	var member models.GuildMember
	err := r.repository.GetGuildMemberByGuildIDAndUserID(guildID, userID, &member)

	if err != nil {
		if err := r.repository.GetGuildMemberByUserID(userID, &member); err != nil {
			member.GuildID = guildID
			member.MemberID = userID
			return r.repository.Create(&member)
		}
	}

	return nil
}

func (r *GuildMemberService) IncrementMessageCount(guildID, userID string) {
	var member models.GuildMember

	if err := r.repository.GetGuildMemberByGuildIDAndUserID(guildID, userID, &member); err != nil {
		slog.Error("Error getting member: %v", err)
		return
	}

	member.MessageCount++
	if err := r.repository.Update(&member); err != nil {
		slog.Error("Error updating member: %v", err)
	}
}

func (r *GuildMemberService) IncrementCommandCount(guildID, userID string) {
	var member models.GuildMember

	if err := r.repository.GetGuildMemberByGuildIDAndUserID(guildID, userID, &member); err != nil {
		slog.Error("Error getting member: %v", err)
		return
	}

	member.CommandCount++
	if err := r.repository.Update(&member); err != nil {
		slog.Error("Error updating member: %v", err)
	}
}

func (r *GuildMemberService) GetBalance(guildID, userID string) (uint, error) {
	var member models.GuildMember

	if err := r.repository.GetGuildMemberByGuildIDAndUserID(guildID, userID, &member); err != nil {
		return 0, err
	}

	return member.Coins, nil
}
