package bot

import (
	"github.com/diamondburned/arikawa/v3/discord"
	"github.com/mauriciofsnts/vulcano/internal/config"
	"github.com/mauriciofsnts/vulcano/internal/discord/bot"
	"github.com/mauriciofsnts/vulcano/internal/discord/t"
	"github.com/mauriciofsnts/vulcano/internal/i18n"
)

func init() {
	bot.RegisterCommand("help", bot.Command{
		Name:        "help",
		Aliases:     []string{"help"},
		Description: t.Translate().Commands.Help.Description.Str(),
		Category:    bot.CategoryBot,
		Handler: func(ctx *bot.Context) {
			categories := bot.GetCategories()

			var fields []discord.EmbedField

			for category, commands := range categories {
				var value string

				for _, command := range commands {
					value += buildCommandResponse(command)
				}

				fields = append(fields, discord.EmbedField{
					Name:  category,
					Value: value,
				})
			}

			ctx.Reply(bot.ComplexMessageData{
				Embed: discord.Embed{
					Title:       t.Translate().Commands.Help.Title.Str(),
					Description: i18n.Replace(t.Translate().Commands.Help.Response.Str(), config.Vulcano.Prefix),
					Fields:      fields,
				},
			})

		},
	})
}

func buildCommandResponse(command bot.Command) string {
	var response string

	response += `****` + config.Vulcano.Prefix + command.Name
	response += "(" + buildAliasesResponse(command.Aliases) + ") ****"

	if command.Description != "" {
		response += " - " + command.Description
	}

	response += "\n"

	return response
}

func buildAliasesResponse(aliases []string) string {
	var response string

	for _, alias := range aliases {
		if response != "" {
			response += ", "
		}

		response += alias
	}

	return response
}
