package code

import (
	"fmt"

	"github.com/diamondburned/arikawa/v3/discord"
	"github.com/lucasb-eyer/go-colorful"
	"github.com/mauriciofsnts/vulcano/internal/config"
	"github.com/mauriciofsnts/vulcano/internal/discord/bot"
)

func init() {
	bot.RegisterCommand("color", bot.Command{
		Name:        "color",
		Description: "Generate a bunch of color stuff based on a given one.",
		Category:    bot.CategoryCode,
		Aliases:     []string{"color"},
		Parameters: []discord.CommandOption{
			&discord.StringOption{
				OptionName:  "color",
				Description: "RGB or HEX color to generate",
				Required:    true,
			},
		},
		Handler: func(ctx *bot.Context) {
			args := ctx.RawArgs

			if len(args) == 0 {
				ctx.ReplyError(bot.ComplexMessageData{
					Embed: discord.Embed{
						Title:       "Color",
						Description: fmt.Sprintf("You need to provide a color to generate. Example: `%scolor #ffffff`", config.Vulcano.Prefix),
					},
				})
				return
			}

			color, err := hexToColor(args[0])

			if err != nil {
				ctx.ReplyError(bot.ComplexMessageData{
					Embed: discord.Embed{
						Title:       "Color",
						Description: fmt.Sprintf("Invalid color provided. Example: `%scolor #ffffff`", config.Vulcano.Prefix),
					},
				})
			}

			r, g, b := color.RGB255()
			Hhsl, Shsl, Lhsl := color.Hsl()
			Hhsv, Shsv, Vhsv := color.Hsv()

			ctx.Reply(bot.ComplexMessageData{
				Embed: discord.Embed{
					Title:       "Color",
					Description: "Color generated successfully!",
					Fields: []discord.EmbedField{
						{
							Name:   "RGBA",
							Value:  fmt.Sprintf("```R: %d\nG: %d\nB: %d\n```", r, g, b),
							Inline: true,
						},
						{
							Name:   "HSV",
							Value:  fmt.Sprintf("```H: %.2f\nS: %.2f\nV: %.2f\n```", Hhsl, Shsl, Vhsv),
							Inline: true,
						},
						{
							Name:   "HSL",
							Value:  fmt.Sprintf("```H: %.2f\nS: %.2f\nL: %.2f\n```", Hhsv, Shsv, Lhsl),
							Inline: true,
						},
					},
				},
			})

		},
	})
}

func hexToColor(hex string) (*colorful.Color, error) {
	c, err := colorful.Hex(hex)
	if err != nil {
		return nil, err
	}
	return &c, nil
}
