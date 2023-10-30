package utils

import (
	"github.com/bwmarrin/discordgo"
	"github.com/google/uuid"
	"github.com/mauriciofsnts/vulcano/internal/discord/events"
	"github.com/mauriciofsnts/vulcano/internal/helpers"
)

func init() {
	events.Register("generate", events.CommandInfo{
		Function: func(cm events.CommandMessage) {
			var args = cm.GetArgString(0)

			switch args {
			case "uuid":
				uuid, err := uuid.NewUUID()
				if err != nil {
					return
				}
				cm.Ok(&discordgo.MessageEmbed{Description: uuid.String()})
			case "cpf":
				cpf, _ := helpers.GenerateCPF()

				cm.Ok(&discordgo.MessageEmbed{Description: cpf})

			default:
				cm.Error(&discordgo.MessageEmbed{Description: "Opção inválida."})
			}

		},
		ApplicationCommand: &discordgo.ApplicationCommand{
			Name:        "generate",
			Description: "Generate various useful information for developers",
			Options: []*discordgo.ApplicationCommandOption{
				{
					Name:        "option",
					Description: "Type of information you want to generate",
					Required:    true,
					Type:        discordgo.ApplicationCommandOptionString,
					Choices: []*discordgo.ApplicationCommandOptionChoice{
						{
							Name:  "uuid",
							Value: "uuid",
						},
						{
							Name:  "cpf",
							Value: "cpf",
						},
					},
				},
			},
		},
	})
}
