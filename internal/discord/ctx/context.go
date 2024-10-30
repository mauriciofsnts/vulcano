package ctx

import (
	"time"

	"github.com/disgoorg/disgo/bot"
	"github.com/disgoorg/disgo/discord"
	"github.com/disgoorg/snowflake/v2"
)

type EventType string

const (
	SLASH_COMMAND EventType = "SLASH_COMMAND"
	MESSAGE       EventType = "MESSAGE"
	BTN_COMPONENT EventType = "BTN_COMPONENT"
)

type TriggerEvent struct {
	AuthorId       snowflake.ID
	ChannelId      snowflake.ID
	GuildId        snowflake.ID
	MessageId      snowflake.ID
	EventTimestamp time.Time
}

type Context struct {
	BotStartAt   time.Time
	TriggerEvent TriggerEvent
	Client       bot.Client
	Type         EventType
	Args         []string
	Response     Response
}

type Response struct {
	Reply    func(title string, description string, fields []discord.EmbedField) discord.MessageCreate
	ReplyErr func(err error) discord.MessageCreate
}

func Execute(
	args []string,
	command *Command,
	trigger TriggerEvent,
	eventType EventType,
	botStartAt time.Time,
	client bot.Client,
) *discord.MessageCreate {

	ctx := Context{
		TriggerEvent: trigger,
		Type:         eventType,
		Args:         args,
		BotStartAt:   botStartAt,
		Client:       client,
		Response:     Response{Reply: Reply, ReplyErr: ReplyErr},
	}

	return command.Handler(ctx)
}
