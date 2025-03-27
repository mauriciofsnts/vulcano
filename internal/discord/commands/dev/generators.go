package dev

import (
	"errors"
	"fmt"
	"regexp"
	"strings"

	"github.com/disgoorg/disgo/discord"
	"github.com/google/uuid"
	"github.com/mauriciofsnts/bot/internal/discord/ctx"
	"github.com/mauriciofsnts/bot/internal/i18n"
	"github.com/mauriciofsnts/bot/internal/utils"
)

func init() {
	ctx.RegisterCommand("generate", ctx.Command{
		Name:        "generate",
		Aliases:     []string{"gen", "g", "cpf", "uuid", "cnpj"},
		Description: ctx.T().Commands.Generate.Description.Str(),
		Options: []discord.ApplicationCommandOption{
			discord.ApplicationCommandOptionString{
				Name:        "option",
				Description: ctx.T().Commands.Generate.Options.Str(),
				Required:    true,
				Choices: []discord.ApplicationCommandOptionChoiceString{
					{Name: "CPF", Value: "cpf"},
					{Name: "CNPJ", Value: "cnpj"},
					{Name: "UUID", Value: "uuid"},
				},
			},
		},
		Handler: func(data ctx.CommandExecutionContext) *discord.MessageCreate {
			option := determineOption(data)

			var generatedContent string

			switch option {
			case "cnpj":
				generatedContent = generateCNPJ(ctx.T())
			case "cpf":
				generatedContent = generateCPF(ctx.T())
			case "uuid":
				generatedContent = generateUUID()

			default:
				return buildErrorResponse(data, string(ctx.T().Commands.Generate.ParamError))
			}

			description := ctx.T().Commands.Generate.Warning.Str() + generatedContent

			msg := i18n.Replace(strings.ToUpper(ctx.T().Commands.Generate.Reply.Str()), option)
			reply := data.Response.BuildDefaultEmbedMessage(msg, description, nil)

			return &reply
		},
	})
}

func determineOption(data ctx.CommandExecutionContext) string {
	if len(data.Args) > 0 {
		return data.Args[0]
	}

	if data.TriggerEvent.TriggeredAlias != "" {
		return data.TriggerEvent.TriggeredAlias
	}

	return ""
}

func buildErrorResponse(data ctx.CommandExecutionContext, message string) *discord.MessageCreate {
	reply := data.Response.BuildDefaultErrorMessage(errors.New(message))
	return &reply
}

func generateCNPJ(t i18n.Language) string {
	cnpj := utils.GenerateCNPJ()
	cnpjMaskRe := regexp.MustCompile(`^(\d{2})(\d{3})(\d{3})(\d{4})(\d{2})$`)
	components := cnpjMaskRe.FindStringSubmatch(cnpj)
	maskedCNPJ := fmt.Sprintf(
		"%s.%s.%s/%s-%s",
		components[1], components[2], components[3], components[4], components[5],
	)
	return i18n.Replace(t.Commands.Generate.WithMask.Str(), cnpj, maskedCNPJ)
}

func generateCPF(t i18n.Language) string {
	cpf := utils.GenerateCPF()
	cpfMaskRe := regexp.MustCompile(`^(\d{3})(\d{3})(\d{3})(\d{2})$`)
	components := cpfMaskRe.FindStringSubmatch(cpf)
	maskedCPF := fmt.Sprintf(
		"%s.%s.%s-%s",
		components[1], components[2], components[3], components[4],
	)
	return i18n.Replace(t.Commands.Generate.WithMask.Str(), cpf, maskedCPF)
}

func generateUUID() string {
	uuid, _ := uuid.NewUUID()
	return fmt.Sprintf("```%s```", uuid)
}
