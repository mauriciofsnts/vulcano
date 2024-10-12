package ctx

import (
	"log/slog"
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
	slog.Info("Trigger event: ", slog.Any("trigger", trigger))

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
