package events

import (
	"fmt"
	"log/slog"
	"math/rand"
	"time"

	"github.com/disgoorg/disgo/bot"
	"github.com/disgoorg/disgo/discord"
	"github.com/mauriciofsnts/bot/internal/utils"

	discordEvents "github.com/disgoorg/disgo/events"
)

func OnGamble(event *discordEvents.MessageReactionAdd, client bot.Client) {
	if event.Member.User.Bot || *event.Emoji.Name != "🎲" || event.GuildID == nil {
		return
	}

	// roll a dice from 1 to 9
	roll := rand.Intn(9) + 1

	if roll == 1 {
		client.Rest().CreateMessage(event.ChannelID, discord.MessageCreate{
			Content: "You rolled a 1, you win!",
		})

		err := client.Rest().DeleteMessage(event.ChannelID, event.MessageID)

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

		go func() {
			time.Sleep(2 * time.Minute)

			client.Rest().UpdateMember(*event.GuildID, event.Member.User.ID, discord.MemberUpdate{
				Mute: utils.PtrTo(false),
			})

		}()
	}
}
