package utils

import (
	"encoding/json"
	"io"
	"os"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/mauriciofsnts/vulcano/internal/discord/events"
	"github.com/pauloo27/logger"
)

type Holiday struct {
	Name string `json:"name"`
	Date string `json:"date"`
	Type string `json:"type"`
}

func init() {
	events.Register("holiday", events.CommandInfo{
		Function: func(cm events.CommandMessage) {
			today := time.Now()

			jsonFile, err := os.Open("./internal/providers/holiday/dates.json")

			if err != nil {
				logger.Debug("Cannot find dates.json file", err.Error())
			}

			defer jsonFile.Close()

			byteValue, err := io.ReadAll(jsonFile)

			if err != nil {
				logger.Debug("Error on read json file", err.Error())
			}

			var holidays []Holiday

			err = json.Unmarshal(byteValue, &holidays)

			if err != nil {
				logger.Debug("Error on unmarshal", err.Error())
			}

			for _, holiday := range holidays {
				holidayDate, err := time.Parse("02/01", holiday.Date)

				if err != nil {
					logger.Debug("Error on parse date", err.Error())
				}

				holidayDate = holidayDate.AddDate(today.Year(), 0, 0)

				if holidayDate.After(today) || holidayDate.Equal(today) {
					var description string

					if holidayDate.Equal(today) {
						description = "O feriado de hoje é " + holiday.Name + "." + "\n" + "Data: " + holiday.Date + "." + "\n" + "Tipo: " + holiday.Type + "."
					} else {
						description = "O próximo feriado é " + holiday.Name + "." + "\n" + "Data: " + holiday.Date + "." + "\n" + "Tipo: " + holiday.Type + "."
					}

					cm.Ok(&discordgo.MessageEmbed{
						Title:       cm.T.Commands.Holiday.Title.Str(),
						Description: description,
					})
					break
				}
			}

		},
		ApplicationCommand: &discordgo.ApplicationCommand{
			Name:        "holiday",
			Description: "Returns the next brazillian holiday",
		},
	})
}
