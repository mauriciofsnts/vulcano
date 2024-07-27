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
	Build          func(embed discord.Embed) discord.MessageCreate
	Embed          func(title string, description string, fields []discord.EmbedField) discord.Embed
	ErrorEmbed     func(err error) discord.Embed
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
		Build:          Build,
		Embed:          Embed,
		ErrorEmbed:     ErrorEmbed,
		BotStartAt:     botStartAt,
		Client:         client,
	}

	return command.Handler(ctx)
}

func Build(embed discord.Embed) discord.MessageCreate {
	builder := discord.NewMessageCreateBuilder()
	builder.SetEmbeds(embed)
	return builder.Build()
}

func Embed(
	title string,
	description string,
	fields []discord.EmbedField,
) discord.Embed {
	embedBuilder := discord.NewEmbedBuilder()
	embedBuilder.
		SetTitle(title).
		SetDescription(description).
		SetColor(0xffffff).
		SetFields(fields...)

	return embedBuilder.Build()
}

func ErrorEmbed(
	err error,
) discord.Embed {
	embedBuilder := discord.NewEmbedBuilder()
	embedBuilder.
		SetTitle("An error occurred").
		SetDescription(err.Error()).
		SetColor(0xff0000)

	return embedBuilder.Build()
}
