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
		Parameters: []discord.CommandOption{
			&discord.StringOption{
				OptionName:  "command",
				Description: "Command to get help",
				Required:    false,
			},
		},
		Handler: func(ctx *bot.Context) {
			args := ctx.RawArgs

			if len(args) > 0 {
				cmd, found := bot.GetCommand(args[0])

				if !found {
					ctx.ReplyError(bot.ComplexMessageData{
						Embed: discord.Embed{
							Title:       "Not found",
							Description: "Command not found",
						},
					})

					return
				}

				commandEmbed := buildHelpCommandEmbed(*cmd)

				ctx.Reply(bot.ComplexMessageData{Embed: commandEmbed})

				return
			}

			embed := buildListCommandsEmbed()

			ctx.Reply(bot.ComplexMessageData{Embed: embed})

		},
	})
}

func buildListCommandsEmbed() discord.Embed {
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

	return discord.Embed{
		Title:       t.Translate().Commands.Help.Title.Str(),
		Description: i18n.Replace(t.Translate().Commands.Help.Response.Str(), config.Vulcano.Prefix),
		Fields:      fields,
	}
}

func buildHelpCommandEmbed(command bot.Command) discord.Embed {
	return discord.Embed{
		Title:       command.Name,
		Description: command.Description,
		Fields: []discord.EmbedField{
			{
				Name:   t.Translate().Utils.Usage.Str(),
				Value:  buildUsageResponse(command),
				Inline: true,
			},
			{
				Name:   t.Translate().Utils.Aliases.Str(),
				Value:  buildAliasesResponse(command.Aliases),
				Inline: true,
			},
		},
	}
}

func buildUsageResponse(command bot.Command) string {
	var response string

	response += config.Vulcano.Prefix + command.Name

	for _, param := range command.Parameters {
		response += " " + param.Name()
	}

	return response
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
