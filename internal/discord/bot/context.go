package bot

import (
	"time"

	"github.com/diamondburned/arikawa/v3/discord"
	"github.com/mauriciofsnts/vulcano/internal/i18n"
)

type TriggerType string

type TriggerEvent struct {
	EventTime     time.Time
	Type          TriggerType
	MessageID     *discord.MessageID
	InteractionID *discord.InteractionID
	Token         string
	GuildID       string
	AuthorID      string
	ChannelID     discord.ChannelID
}

const (
	SuccessColor = 0xffffff
	ErrorColor   = 0xff5555
)

type Context struct {
	startTime         time.Time
	Args              []any
	RawArgs           []string
	Type              TriggerType
	Bot               Bot
	AuthorID, GuildID string
	MessageID         discord.MessageID
	Command           *Command
	TriggerType       TriggerType
	T                 i18n.Language
}

func (ctx *Context) SuccessEmbed(embed discord.Embed) discord.Embed {
	embed.Color = SuccessColor
	return embed
}

func (ctx *Context) ErrorEmbed(embed discord.Embed) discord.Embed {
	embed.Color = ErrorColor
	return embed
}
