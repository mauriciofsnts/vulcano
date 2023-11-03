package utils

import (
	"encoding/json"
	"io"
	"os"
	"time"

	"github.com/diamondburned/arikawa/v3/discord"
	"github.com/mauriciofsnts/vulcano/internal/discord/bot"
	"github.com/pauloo27/logger"
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
			Name:    "holiday",
			Aliases: []string{"feriado"},
			Handler: func(ctx *bot.Context) discord.Embed {
				jsonFile, err := os.Open("./internal/providers/holiday/dates.json")

				if err != nil {
					logger.Debug("Cannot find dates.json file", err.Error())
					return ctx.ErrorEmbed(discord.Embed{
						Title:       "Feriados",
						Description: "Não foi possível encontrar o arquivo de feriados.",
					})
				}

				defer jsonFile.Close()

				byteValue, err := io.ReadAll(jsonFile)

				if err != nil {
					logger.Debug("Cannot read dates.json file", err.Error())
					return ctx.ErrorEmbed(discord.Embed{
						Title:       "Feriados",
						Description: "Não foi possível ler o arquivo de feriados.",
					})
				}

				var holidays []Holiday

				err = json.Unmarshal(byteValue, &holidays)

				if err != nil {
					logger.Debug("Cannot unmarshal dates.json file", err.Error())
					return ctx.ErrorEmbed(discord.Embed{
						Title:       "Feriados",
						Description: "Não foi possível ler o arquivo de feriados.",
					})
				}

				today := time.Now()
				var description string

				for _, holiday := range holidays {
					holidayDate, err := time.Parse("02/01", holiday.Date)

					if err != nil {
						logger.Debug("Cannot parse holiday date", err.Error())
						return ctx.ErrorEmbed(discord.Embed{
							Title:       "Feriados",
							Description: "Não foi possível ler o arquivo de feriados.",
						})
					}

					holidayDate = holidayDate.AddDate(today.Year(), 0, 0)

					if holidayDate.After(today) || holidayDate.Equal(today) {
						if holidayDate.Equal(today) {
							description = "O feriado de hoje é " + holiday.Name + "." + "\n" + "Data: " + holiday.Date + "." + "\n" + "Tipo: " + holiday.Type + "."
						} else {
							description = "O próximo feriado é " + holiday.Name + "." + "\n" + "Data: " + holiday.Date + "." + "\n" + "Tipo: " + holiday.Type + "."
						}
					}
				}

				embed := discord.Embed{
					Title:       "TabNews",
					Description: description,
				}

				return ctx.SuccessEmbed(embed)
			},
		},
	)
}
