package events

import (
	"fmt"
	"math/rand"

	"github.com/disgoorg/disgo/bot"
	"github.com/disgoorg/disgo/discord"

	discordEvents "github.com/disgoorg/disgo/events"
)

func OnGamble(event *discordEvents.MessageReactionAdd, client bot.Client) {
	if event.Member.User.Bot || *event.Emoji.Name != "🎲" {
		return
	}

	if client == nil {
		return
	}

	// roll a dice from 1 to 10
	roll := rand.Intn(10) + 1

	if roll == 1 {
		client.Rest().CreateMessage(event.ChannelID, discord.MessageCreate{
			Content: "You rolled a 1, you lose!",
		})
	} else {
		client.Rest().CreateMessage(event.ChannelID, discord.MessageCreate{
			Content: "You rolled a " + fmt.Sprint(roll) + ", you win!",
		})
	}
}
