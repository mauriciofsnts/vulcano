package bot

import (
	"time"

	"github.com/diamondburned/arikawa/v3/api"
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
	triggerEvent      TriggerEvent
	T                 i18n.Language
}

func (ctx *Context) Reply(embed discord.Embed) {
	embed.Color = SuccessColor
	ctx.handler([]discord.Embed{embed})
}

func (ctx *Context) ReplyError(embed discord.Embed) {
	embed.Color = ErrorColor
	ctx.handler([]discord.Embed{embed})
}

func (ctx *Context) handler(embeds []discord.Embed) {

	if ctx.TriggerType == CommandTriggerSlash {
		ctx.Bot.State.RespondInteraction(*ctx.triggerEvent.InteractionID, ctx.triggerEvent.Token, api.InteractionResponse{
			Type: api.MessageInteractionWithSource,
			Data: &api.InteractionResponseData{
				Embeds: &embeds,
			},
		})
	}

	if ctx.TriggerType == CommandTriggerMessage {
		ctx.Bot.State.SendEmbedReply(ctx.triggerEvent.ChannelID, *ctx.triggerEvent.MessageID, embeds[0])
	}

}
