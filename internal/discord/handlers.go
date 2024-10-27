package discord

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"strings"

	"github.com/disgoorg/disgo/bot"
	disgo "github.com/disgoorg/disgo/discord"
	"github.com/disgoorg/disgo/events"
	"github.com/disgoorg/snowflake/v2"
	"github.com/mauriciofsnts/bot/internal/config"
	"github.com/mauriciofsnts/bot/internal/discord/ctx"
	customEvents "github.com/mauriciofsnts/bot/internal/discord/events"
	"github.com/mauriciofsnts/bot/internal/providers"
)

func OnMessageCreatedEvent(event *events.MessageCreate, client bot.Client, cfg config.Config) {
	message := event.Message

	if message.Author.Bot {
		return
	}

	if message.GuildID != nil {
		message := event.Message
		guildID := message.GuildID.String()
		authorID := message.Author.ID.String()

		err := providers.Services.GuildMember.EnsureMemberValidity(guildID, authorID)
		if err != nil {
			slog.Error("Error ensuring member validity: ", err)
		}

		if !strings.HasPrefix(message.Content, cfg.Discord.Prefix) {
			providers.Services.GuildMember.IncrementMessageCount(guildID, authorID)
			return
		}

		executeMessageCommand(event, client, cfg, true)
	} else {
		if !strings.HasPrefix(event.Message.Content, cfg.Discord.Prefix) {
			return
		}

		executeMessageCommand(event, client, cfg, false)
	}
}

func OnInteractionCreatedEvent(event *events.ApplicationCommandInteractionCreate, client bot.Client) {
	data := event.SlashCommandInteractionData()

	commandName := data.CommandName()
	found, cmd := ctx.GetCommandByAlias(commandName)

	if !found {
		slog.Error("Command not found: ", slog.String("command", commandName))
		slog.Error("The user can't dispatch an unknown slash command. Check if the command is registered")
		return
	}

	trigger := ctx.TriggerEvent{
		ChannelId:      event.Channel().ID(),
		EventTimestamp: event.CreatedAt(),
		AuthorId:       event.User().ID,
	}

	guildId := event.GuildID()

	if guildId != nil {
		trigger.GuildId = *guildId

		err := providers.Services.GuildMember.EnsureMemberValidity(guildId.String(), event.User().ID.String())

		if err != nil {
			slog.Error("Error ensuring member validity: ", err)
		}
	}

	var args []string

	for _, option := range data.Options {
		var value any

		err := json.Unmarshal(option.Value, &value)

		if err != nil {
			return
		}

		args = append(args, fmt.Sprintf("%v", value))
	}

	msg := ctx.Execute(args, cmd, trigger, ctx.SLASH_COMMAND, StartedAt, client)

	if msg != nil {
		event.CreateMessage(*msg)
		providers.Services.GuildMember.IncrementCommandCount(guildId.String(), event.User().ID.String())
	}
}

func OnReadyEvent(event *events.Ready) {
	slog.Info("Bot is ready!")
}

func OnGuildChannelCreatedEvent(event *events.GuildChannelCreate, client bot.Client) {
	channelId := event.ChannelID
	message := disgo.NewMessageCreateBuilder().SetContent("first!").Build()
	event.Client().Rest().CreateMessage(channelId, message)
}

func OnMessageReactionAddedEvent(event *events.MessageReactionAdd, client bot.Client) {
	customEvents.OnGamble(event, client)
}

func OnComponentInteractionEvent(event *events.ComponentInteractionCreate, client bot.Client) {
	found, componentState := ctx.GetComponentStateInDatabase(event.Message.ID.String())

	if !found {
		slog.Error("Button state not found: ", slog.String("message id", event.Message.ID.String()))
		return
	}

	trigger := ctx.TriggerEvent{
		AuthorId:       snowflake.MustParse(componentState.AuthorID),
		ChannelId:      snowflake.MustParse(componentState.ChannelID),
		GuildId:        snowflake.MustParse(componentState.GuildID),
		MessageId:      snowflake.MustParse(componentState.MessageID),
		EventTimestamp: event.CreatedAt(),
	}

	state := ctx.ComponentState{
		TriggerEvent: trigger,
		Client:       client,
		State:        componentState.State,
	}

	found, handler := ctx.GetComponentHandlerByName(componentState.Command)

	if !found {
		slog.Error("Component handler not found: ", slog.String("command", componentState.Command))
		return
	}

	handler(event, &state)
}

func OnGuildReady(event *events.GuildReady) {
	providers.Services.Guild.EnsureGuildExists(event.Guild)
}

func executeMessageCommand(event *events.MessageCreate, client bot.Client, cfg config.Config, isGuild bool) {
	message := event.Message
	inputMessage := strings.Split(message.Content, " ")
	commandName := strings.TrimPrefix(inputMessage[0], cfg.Discord.Prefix)
	found, cmd := ctx.GetCommandByAlias(commandName)

	if !found {
		return
	}

	args := inputMessage[1:]

	trigger := ctx.TriggerEvent{
		AuthorId:       message.Author.ID,
		ChannelId:      message.ChannelID,
		MessageId:      message.ID,
		EventTimestamp: message.CreatedAt,
	}

	if isGuild {
		trigger.GuildId = *message.GuildID
	}

	msg := ctx.Execute(args, cmd, trigger, ctx.MESSAGE, StartedAt, client)

	if msg != nil {
		msg.MessageReference = &disgo.MessageReference{MessageID: &message.ID}
		event.Client().Rest().CreateMessage(event.ChannelID, *msg)

		if isGuild {
			providers.Services.GuildMember.IncrementCommandCount(message.GuildID.String(), message.Author.ID.String())
		}
	}
}
