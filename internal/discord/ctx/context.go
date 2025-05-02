package ctx

import (
	"time"

	"github.com/disgoorg/disgo/bot"
	"github.com/disgoorg/disgo/discord"
	"github.com/disgoorg/snowflake/v2"
	"github.com/mauriciofsnts/bot/internal/config"
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
	TriggeredAlias string
}

type CommandExecutionContext struct {
	BotStartAt   time.Time
	TriggerEvent TriggerEvent
	Client       bot.Client
	Type         EventType
	Args         []string
	Response     Response
	Config       config.Config
	Cache        Cache
}

type Cache struct {
	CreateCommandState func(id string, data GuildState) bool
	GetCommandState    func(id string) (GuildState, bool)
	DeleteCommandState func(id string) bool
}

type Response struct {
	BuildDefaultEmbedMessage func(title string, description string, fields []discord.EmbedField) discord.MessageCreate
	BuildDefaultErrorMessage func(err error) discord.MessageCreate
}

func Execute(
	args []string,
	command *Command,
	trigger TriggerEvent,
	eventType EventType,
	botStartAt time.Time,
	client bot.Client,
	cfg config.Config,
) *discord.MessageCreate {

	ctx := CommandExecutionContext{
		TriggerEvent: trigger,
		Type:         eventType,
		Args:         args,
		BotStartAt:   botStartAt,
		Client:       client,
		Response:     Response{BuildDefaultEmbedMessage: BuildDefaultEmbedMessage, BuildDefaultErrorMessage: BuildDefaultErrorMessage},
		Cache: Cache{
			GetCommandState:    GetCommandState,
			DeleteCommandState: DeleteCommandState,
			CreateCommandState: func(id string, data GuildState) bool {
				return CreateCommandState(id, data, 24*time.Hour)
			},
		},
		Config: cfg,
	}

	return command.Handler(ctx)
}
