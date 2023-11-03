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
	Bot          Bot
	T            i18n.Language
	startTime    time.Time
	Command      *Command
	RawArgs      []string
	AuthorID     string
	GuildID      string
	MessageID    discord.MessageID
	TriggerType  TriggerType
	triggerEvent TriggerEvent
}

func (ctx *Context) Reply(embed discord.Embed) {
	embed.Color = SuccessColor
	embed = ctx.appendExecutionInfoToEmbed(embed)
	ctx.handleReply([]discord.Embed{embed})
}

func (ctx *Context) ReplyError(embed discord.Embed) {
	embed.Color = ErrorColor
	embed = ctx.appendExecutionInfoToEmbed(embed)
	ctx.handleReply([]discord.Embed{embed})
}

func (ctx *Context) handleReply(embeds []discord.Embed) {
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

func (ctx *Context) appendExecutionInfoToEmbed(embed discord.Embed) discord.Embed {
	took := "Took: " + time.Since(ctx.startTime).Truncate(time.Second).String()

	embed.Footer = &discord.EmbedFooter{
		Text: took,
	}

	return embed
}
