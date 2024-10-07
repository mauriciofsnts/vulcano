package events

import (
	"fmt"
	"log/slog"
	"math/rand"

	"github.com/disgoorg/disgo/bot"
	"github.com/disgoorg/disgo/discord"
	"github.com/mauriciofsnts/bot/internal/utils"

	discordEvents "github.com/disgoorg/disgo/events"
)

func OnGamble(event *discordEvents.MessageReactionAdd, client bot.Client) {
	// this command is only for users and not bots
	// and the event must be in a guild
	if event.Member.User.Bot || *event.Emoji.Name != "🎲" || event.GuildID == nil {
		return
	}

	// roll a dice from 1 to 10
	roll := rand.Intn(10) + 1

	if roll == 1 {
		client.Rest().CreateMessage(event.ChannelID, discord.MessageCreate{
			Content: "You rolled a 1, you win!",
		})

		_, err := client.Rest().UpdateMember(*event.GuildID, event.Member.User.ID, discord.MemberUpdate{
			Mute: utils.PtrTo(false),
		})

		if err != nil {
			slog.Error(err.Error())
		}
	} else {
		client.Rest().CreateMessage(event.ChannelID, discord.MessageCreate{
			Content: "You rolled a " + fmt.Sprint(roll) + ", you lose!",
		})

		_, err := client.Rest().UpdateMember(*event.GuildID, event.Member.User.ID, discord.MemberUpdate{
			Mute: utils.PtrTo(true),
		})

		if err != nil {
			slog.Error(err.Error())
		}
	}
}
