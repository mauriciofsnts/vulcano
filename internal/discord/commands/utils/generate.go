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
			var args string

			if cm.Interaction != nil {
				args = cm.Interaction.Args[0].StringValue()
			} else {
				args = cm.Message.Args[0]
			}

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
			Description: "Gerar várias informações uteis para desenvolvedores.",
			Options: []*discordgo.ApplicationCommandOption{
				{
					Name:        "option",
					Description: "Tipo de informação que deseja gerar",
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
