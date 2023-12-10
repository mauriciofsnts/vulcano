package utils

import (
	"encoding/json"
	"io"
	"os"
	"time"

	"github.com/diamondburned/arikawa/v3/discord"
	"github.com/mauriciofsnts/vulcano/internal/discord/bot"
)

type Holiday struct {
	Name string `json:"name"`
	Date string `json:"date"`
	Type string `json:"type"`
}

func init() {
	bot.RegisterCommand(
		"holiday",
		bot.Command{
			Name:     "holiday",
			Aliases:  []string{"feriado"},
			Category: "🔧 Utils",
			Handler: func(ctx *bot.Context) {
				commandTitle := ctx.T.Commands.Holiday.Title.Str()
				jsonFile, err := os.Open("./internal/providers/holiday/dates.json")

				if err != nil {
					ctx.ReplyError(bot.ComplexMessageData{
						Embed: discord.Embed{
							Title:       commandTitle,
							Description: "It was not possible to find the holiday file",
						},
					})
					return
				}

				defer jsonFile.Close()

				byteValue, err := io.ReadAll(jsonFile)

				if err != nil {
					ctx.ReplyError(bot.ComplexMessageData{
						Embed: discord.Embed{
							Title:       commandTitle,
							Description: "Unable to read the holidays file",
						},
					})
					return
				}

				var holidays []Holiday

				err = json.Unmarshal(byteValue, &holidays)

				if err != nil {
					ctx.ReplyError(bot.ComplexMessageData{
						Embed: discord.Embed{
							Title:       commandTitle,
							Description: "Unable to read the holidays file",
						},
					})
					return
				}

				today := time.Now()
				var description string

				for _, holiday := range holidays {
					holidayDate, err := time.Parse("02/01", holiday.Date)

					if err != nil {
						ctx.ReplyError(bot.ComplexMessageData{
							Embed: discord.Embed{
								Title:       commandTitle,
								Description: "Unable to read the holidays file",
							},
						})
						return
					}

					holidayDate = holidayDate.AddDate(today.Year(), 0, 0)

					if holidayDate.After(today) || holidayDate.Equal(today) {
						if holidayDate.Equal(today) {
							description = "Today is " + holiday.Name + "." + "\n" + "Date: " + holiday.Date + "." + "\n" + "Type: " + holiday.Type + "."
						} else {
							description = "The next holiday is " + holiday.Name + "." + "\n" + "Date: " + holiday.Date + "." + "\n" + "Type: " + holiday.Type + "."
						}
					}
				}

				ctx.Reply(bot.ComplexMessageData{
					Embed: discord.Embed{
						Title:       commandTitle,
						Description: description,
					},
				})
			},
		},
	)
}
