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
	successColor = 0xffffff
	errorColor   = 0xff5555
)

type Context struct {
	Bot          Discord
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

type ComplexMessageData struct {
	Embed      discord.Embed
	Components *discord.ContainerComponents
}

func (ctx *Context) Reply(data ComplexMessageData) {
	data.Embed.Color = successColor
	data.Embed = ctx.appendExecutionInfoToEmbed(data.Embed)
	ctx.handleReply(data)

}

func (ctx *Context) ReplyError(data ComplexMessageData) {
	data.Embed.Color = errorColor
	data.Embed = ctx.appendExecutionInfoToEmbed(data.Embed)
	ctx.handleReply(data)
}

func (ctx *Context) handleReply(data ComplexMessageData) {
	if ctx.TriggerType == CommandTriggerSlash {

		response := &api.InteractionResponseData{
			Embeds: &[]discord.Embed{data.Embed},
		}

		if data.Components != nil {
			response.Components = data.Components
		}

		ctx.Bot.State.RespondInteraction(*ctx.triggerEvent.InteractionID, ctx.triggerEvent.Token, api.InteractionResponse{
			Type: api.MessageInteractionWithSource,
			Data: response,
		})
	}

	if ctx.TriggerType == CommandTriggerMessage {

		response := api.SendMessageData{
			Embeds: []discord.Embed{data.Embed},
		}

		if data.Components != nil {
			response.Components = *data.Components
		}

		ctx.Bot.State.SendMessageComplex(ctx.triggerEvent.ChannelID, response)
	}
}

func (ctx *Context) appendExecutionInfoToEmbed(embed discord.Embed) discord.Embed {
	took := "Took: " + time.Since(ctx.startTime).Truncate(time.Second).String()

	embed.Footer = &discord.EmbedFooter{
		Text: took,
	}

	return embed
}
