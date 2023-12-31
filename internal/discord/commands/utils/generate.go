package utils

import (
	"fmt"
	"regexp"

	"github.com/diamondburned/arikawa/v3/discord"
	"github.com/google/uuid"
	"github.com/mauriciofsnts/vulcano/internal/discord/bot"
	"github.com/mauriciofsnts/vulcano/internal/discord/t"
	"github.com/mauriciofsnts/vulcano/internal/i18n"
	"github.com/mauriciofsnts/vulcano/internal/providers/documents"
)

func init() {
	bot.RegisterCommand(
		"generate",
		bot.Command{
			Name:        "generate",
			Description: t.Translate().Commands.Generate.Description.Str(),
			Aliases:     []string{"gen", "g"},
			Category:    bot.CategoryUtils,
			Parameters: []discord.CommandOption{
				&discord.StringOption{
					OptionName:  "type",
					Description: "Type of information to generate. Available types: `cpf`, `uuid`, `cnpj`",
					Choices:     []discord.StringChoice{{Name: "cpf", Value: "cpf"}, {Name: "uuid", Value: "uuid"}, {Name: "cnpj", Value: "cnpj"}},
					Required:    true,
				},
			},
			Handler: func(ctx *bot.Context) {
				var args = ctx.RawArgs

				if len(args) == 0 {
					ctx.ReplyError(bot.ComplexMessageData{
						Embed: discord.Embed{
							Title:       t.Translate().Commands.Generate.Title.Str(),
							Description: i18n.Replace(t.Translate().Errors.MissingParamter.Str(), "tipo"),
						},
					})
					return
				}

				var embed discord.Embed

				switch args[0] {
				case "cnpj":
					cnpj := documents.GenerateCNPJ()

					var cnpjMaskRe = regexp.MustCompile(`^(\d{2})(\d{3})(\d{3})(\d{4})(\d{2})$`)
					components := cnpjMaskRe.FindStringSubmatch(cnpj)

					maskedCNPJ := fmt.Sprintf(
						"%s.%s.%s/%s-%s",
						components[1], components[2], components[3], components[4], components[5],
					)

					description := "With mask: " + cnpj + "\nWithout mask: " + maskedCNPJ

					embed = discord.Embed{
						Title:       t.Translate().Commands.Generate.Title.Str(),
						Description: description,
					}
				case "cpf":
					cpf := documents.GenerateCPF()

					var cpfMaskRe = regexp.MustCompile(`^(\d{3})(\d{3})(\d{3})(\d{2})$`)

					components := cpfMaskRe.FindStringSubmatch(cpf)

					maskedCPF := fmt.Sprintf(
						"%s.%s.%s-%s",
						components[1], components[2], components[3], components[4],
					)

					description := "With mask: " + cpf + "\nWithout mask: " + maskedCPF

					embed = discord.Embed{
						Title:       t.Translate().Commands.Generate.Title.Str(),
						Description: description,
					}
				case "uuid":
					uuid, err := uuid.NewUUID()

					if err != nil {
						return
					}

					embed = discord.Embed{
						Title:       t.Translate().Commands.Generate.Title.Str(),
						Description: uuid.String(),
					}
				default:
					ctx.ReplyError(bot.ComplexMessageData{
						Embed: discord.Embed{
							Title:       t.Translate().Commands.Generate.Title.Str(),
							Description: i18n.Replace(t.Translate().Errors.InvalidParameter.Str(), "tipo"),
						},
					})
					return
				}

				ctx.Reply(bot.ComplexMessageData{
					Embed: embed,
				})
			}})
}
