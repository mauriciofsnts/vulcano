package dev

import (
	"errors"
	"fmt"
	"regexp"

	"github.com/disgoorg/disgo/discord"
	"github.com/google/uuid"
	"github.com/mauriciofsnts/exodia/internal/discord/ctx"
	"github.com/mauriciofsnts/exodia/internal/providers/utils"
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
		Handler: func(ctx *ctx.Context) discord.MessageCreate {
			args := ctx.Args

			if len(args) == 0 {
				return ctx.Reply(
					ctx.ErrorEmbed(
						errors.New("you need to specify the type of information to generate. Available types: `cpf`, `uuid`, `cnpj`"),
					))
			}

			var value string

			switch args[0] {
			case "cnpj":
				cnpj := utils.GenerateCNPJ()

				var cnpjMaskRe = regexp.MustCompile(`^(\d{2})(\d{3})(\d{3})(\d{4})(\d{2})$`)
				components := cnpjMaskRe.FindStringSubmatch(cnpj)

				maskedCNPJ := fmt.Sprintf(
					"%s.%s.%s/%s-%s",
					components[1], components[2], components[3], components[4], components[5],
				)

				value = "With mask: " + cnpj + "\nWithout mask: " + maskedCNPJ

			case "cpf":
				cpf := utils.GenerateCPF()

				var cpfMaskRe = regexp.MustCompile(`^(\d{3})(\d{3})(\d{3})(\d{2})$`)

				components := cpfMaskRe.FindStringSubmatch(cpf)

				maskedCPF := fmt.Sprintf(
					"%s.%s.%s-%s",
					components[1], components[2], components[3], components[4],
				)

				value = "With mask: " + cpf + "\nWithout mask: " + maskedCPF

			case "uuid":
				uuid, _ := uuid.NewUUID()

				value = uuid.String()
			default:
				return ctx.Reply(
					ctx.ErrorEmbed(
						errors.New("invalid type of information to generate. Available types: `cpf`, `uuid`, `cnpj`"),
					))
			}

			return ctx.Reply(ctx.Embed(
				fmt.Sprintf("Generated %s", args[0]),
				value,
				[]discord.EmbedField{},
			))
		},
	})
}
