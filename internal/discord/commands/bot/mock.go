package bot

import (
	"github.com/diamondburned/arikawa/v3/discord"
	"github.com/mauriciofsnts/vulcano/internal/discord/bot"
)

func init() {
	bot.RegisterCommand("mock", bot.Command{
		Name:        "mock",
		Aliases:     []string{"mock"},
		Description: "Just a mock command to test buttons, should be removed later",
		Category:    bot.CategoryBot,
		Handler: func(ctx *bot.Context) {
			ctx.Reply(bot.ComplexMessageData{
				Embed: discord.Embed{
					Title:       "Mock",
					Description: "Mock",
				},
				Components: discord.ComponentsPtr(
					&discord.ActionRowComponent{
						&discord.ButtonComponent{
							Label:    "Hello World!",
							CustomID: "first_button",
							Emoji:    &discord.ComponentEmoji{Name: "👋"},
							Style:    discord.PrimaryButtonStyle(),
						},
					},
				),
			})

		},
	})
}
