package dev

import (
	"errors"
	"fmt"
	"regexp"

	"github.com/disgoorg/disgo/discord"
	"github.com/google/uuid"
	"github.com/mauriciofsnts/bot/internal/discord/ctx"
	"github.com/mauriciofsnts/bot/internal/providers/utils"
)

func init() {
	ctx.AttachCommand("generate", ctx.Command{
		Name:        "generate",
		Aliases:     []string{"gen", "g"},
		Description: "Generate random information",
		Options: []discord.ApplicationCommandOption{
			discord.ApplicationCommandOptionString{
				Name:        "option",
				Description: "Type of information to generate. Available types: `cpf`, `uuid`, `cnpj`",
				Required:    true,
				Choices: []discord.ApplicationCommandOptionChoiceString{
					{Name: "CPF", Value: "cpf"},
					{Name: "CNPJ", Value: "cnpj"},
					{Name: "UUID", Value: "uuid"},
				},
			},
		},
		Handler: generateHandler,
	})
}

func generateHandler(ctx ctx.Context) *discord.MessageCreate {
	args := ctx.Args

	if len(args) == 0 {
		return buildErrorResponse(ctx, "you need to specify the type of information to generate. Available types: `cpf`, `uuid`, `cnpj`")
	}

	var value string

	switch args[0] {
	case "cnpj":
		value = generateCNPJ()
	case "cpf":
		value = generateCPF()
	case "uuid":
		value = generateUUID()
	default:
		return buildErrorResponse(ctx, "invalid type of information to generate. Available types: `cpf`, `uuid`, `cnpj`")
	}

	reply := ctx.Response.Reply(fmt.Sprintf("Generated %s", args[0]), value, nil)
	return &reply
}

func buildErrorResponse(ctx *ctx.Context, message string) *discord.MessageCreate {
	reply := ctx.Response.ReplyErr(errors.New(message))
	return &reply
}

func generateCNPJ() string {
	cnpj := utils.GenerateCNPJ()
	cnpjMaskRe := regexp.MustCompile(`^(\d{2})(\d{3})(\d{3})(\d{4})(\d{2})$`)
	components := cnpjMaskRe.FindStringSubmatch(cnpj)
	maskedCNPJ := fmt.Sprintf(
		"%s.%s.%s/%s-%s",
		components[1], components[2], components[3], components[4], components[5],
	)
	return fmt.Sprintf("With mask: ```%s```\nWithout mask: ```%s```", cnpj, maskedCNPJ)
}

func generateCPF() string {
	cpf := utils.GenerateCPF()
	cpfMaskRe := regexp.MustCompile(`^(\d{3})(\d{3})(\d{3})(\d{2})$`)
	components := cpfMaskRe.FindStringSubmatch(cpf)
	maskedCPF := fmt.Sprintf(
		"%s.%s.%s-%s",
		components[1], components[2], components[3], components[4],
	)
	return fmt.Sprintf("With mask: ```%s```\nWithout mask: ```%s```", cpf, maskedCPF)
}

func generateUUID() string {
	uuid, _ := uuid.NewUUID()
	return uuid.String()
}
