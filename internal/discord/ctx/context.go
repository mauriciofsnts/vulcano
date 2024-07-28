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
)

type TriggerEvent struct {
	AuthorId       snowflake.ID
	ChannelId      snowflake.ID
	GuildId        snowflake.ID
	MessageId      snowflake.ID
	EventTimestamp time.Time
}

type Context struct {
	BotStartAt     time.Time
	CommandStartAt time.Time
	TriggerEvent   TriggerEvent
	Client         bot.Client
	Type           EventType
	Args           []string
	Reply          func(title string, description string, fields []discord.EmbedField) discord.MessageCreate
	ReplyErr       func(err error) discord.MessageCreate
}

func Execute(
	args []string,
	command *Command,
	trigger TriggerEvent,
	eventType EventType,
	botStartAt time.Time,
	client bot.Client,
) *discord.MessageCreate {
	ctx := &Context{
		CommandStartAt: time.Now(),
		TriggerEvent:   trigger,
		Type:           eventType,
		Args:           args,
		Reply:          Reply,
		ReplyErr:       ReplyErr,
		BotStartAt:     botStartAt,
		Client:         client,
	}

	return command.Handler(ctx)
}

func Reply(
	title string,
	description string,
	fields []discord.EmbedField,
) discord.MessageCreate {
	builder := discord.NewMessageCreateBuilder()

	embedBuilder := discord.NewEmbedBuilder()
	embedBuilder.
		SetTitle(title).
		SetDescription(description).
		SetColor(0xffffff).
		SetFields(fields...)
	embed := embedBuilder.Build()

	builder.SetEmbeds(embed)
	return builder.Build()
}

func ReplyErr(
	err error,
) discord.MessageCreate {
	builder := discord.NewMessageCreateBuilder()

	embedBuilder := discord.NewEmbedBuilder()
	embedBuilder.
		SetTitle("An error occurred").
		SetDescription(err.Error()).
		SetColor(0xff0000)

	embed := embedBuilder.Build()
	builder.SetEmbeds(embed)
	return builder.Build()
}
