package utils

import (
	"github.com/diamondburned/arikawa/v3/discord"
	"github.com/google/uuid"
	"github.com/mauriciofsnts/vulcano/internal/discord/bot"
	"github.com/mauriciofsnts/vulcano/internal/providers/documents"
)

func init() {
	bot.RegisterCommand(
		"generate",
		bot.Command{
			Name:        "generate",
			Description: "Generate various useful information for developers",
			Aliases:     []string{"gen", "g"},
			Parameters: []discord.CommandOption{
				&discord.StringOption{
					OptionName:  "type",
					Description: "Type of information to generate. Available types: `cpf`, `uuid`",
					Choices:     []discord.StringChoice{{Name: "cpf", Value: "cpf"}, {Name: "uuid", Value: "uuid"}},
					Required:    true,
				},
			},
			Handler: func(ctx *bot.Context) {
				var args = ctx.RawArgs

				if len(args) == 0 {
					ctx.ReplyError(bot.ComplexMessageData{
						Embed: discord.Embed{
							Title:       "Generate",
							Description: "Você precisa informar o que deseja gerar.",
						},
					})
					return
				}

				var embed discord.Embed

				switch args[0] {
				case "cpf":
					cpf, _ := documents.GenerateCPF()

					embed = discord.Embed{
						Title:       "Generate",
						Description: cpf,
					}
				case "uuid":
					uuid, err := uuid.NewUUID()

					if err != nil {
						return
					}

					embed = discord.Embed{
						Title:       "Generate",
						Description: uuid.String(),
					}
				default:
					ctx.ReplyError(bot.ComplexMessageData{
						Embed: discord.Embed{
							Title:       "Generate",
							Description: "Tipo de informação inválido.",
						},
					})
					return
				}

				ctx.Reply(bot.ComplexMessageData{
					Embed: embed,
				})
			}})
}
