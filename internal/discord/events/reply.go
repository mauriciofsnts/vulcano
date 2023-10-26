package events

import (
	"github.com/bwmarrin/discordgo"
)

func (cmnd CommandMessage) Reply(embed *discordgo.MessageEmbed) {

	isInteraction := cmnd.Interaction != nil

	if isInteraction {
		cmnd.Session.InteractionRespond(
			cmnd.Interaction.Interaction,
			&discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Embeds: []*discordgo.MessageEmbed{embed},
				},
			},
		)
	} else {
		cmnd.Session.ChannelMessageSendEmbed(cmnd.Message.ChannelID, embed)
	}

}

func (cmnd CommandMessage) Error(embed *discordgo.MessageEmbed) {
	embed.Color = 0xe33e32
	cmnd.Reply(embed)
}

func (cmnd CommandMessage) Ok(embed *discordgo.MessageEmbed) {
	embed.Color = 0x42f54b
	cmnd.Reply(embed)
}
