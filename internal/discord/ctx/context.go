package ctx

import (
	"time"

	"github.com/disgoorg/disgo/discord"
)

type EventType string

const (
	SLASH_COMMAND EventType = "SLASH_COMMAND"
	MESSAGE       EventType = "MESSAGE"
)

type TriggerEvent struct {
	AuthorId       string
	ChannelId      string
	GuildId        string
	MessageId      string
	EventTimestamp time.Time
}

type Context struct {
	StartTimestamp time.Time
	TriggerEvent   TriggerEvent
	Type           EventType
	Args           []string
	Reply          func(embed discord.Embed) discord.MessageCreate
	Embed          func(title string, description string, fields []discord.EmbedField) discord.Embed
	ErrorEmbed     func(err error) discord.Embed
}

func Execute(
	args []string,
	command *Command,
	trigger TriggerEvent,
	eventType EventType,
) discord.MessageCreate {
	ctx := &Context{
		StartTimestamp: time.Now(),
		TriggerEvent:   trigger,
		Type:           eventType,
		Args:           args,
		Reply:          Reply,
		Embed:          Embed,
		ErrorEmbed:     ErrorEmbed,
	}

	return command.Handler(ctx)
}

func Reply(embed discord.Embed) discord.MessageCreate {
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
