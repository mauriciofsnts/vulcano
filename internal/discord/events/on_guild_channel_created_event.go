package events

import (
	"github.com/disgoorg/disgo/bot"
	disgo "github.com/disgoorg/disgo/discord"
	disgoEvents "github.com/disgoorg/disgo/events"
)

func OnGuildChannelCreatedEvent(event *disgoEvents.GuildChannelCreate, client bot.Client) {
	channelId := event.ChannelID
	message := disgo.NewMessageCreateBuilder().SetContent("first!").Build()
	event.Client().Rest().CreateMessage(channelId, message)
}
