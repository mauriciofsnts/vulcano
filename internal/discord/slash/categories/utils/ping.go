package utils

import (
	"github.com/bwmarrin/discordgo"
	"github.com/mauriciofsnts/vulcano/internal/discord/slash"
	"github.com/mauriciofsnts/vulcano/internal/i18n"
)

func init() {
	slash.RegisterSlashCommand(
		&slash.SlashCommand{
			ApplicationCommand: &discordgo.ApplicationCommand{
				Name:        "ping",
				Description: "Ping the bot",
			},
			Handler: func(ctx *slash.DiscordContext, t *i18n.Language) {

				ctx.Ok(&discordgo.MessageEmbed{
					Title:       t.Commands.Ping.Title.Str(),
					Description: t.Commands.Ping.Response.Str(),
				})

			},
		})
}
