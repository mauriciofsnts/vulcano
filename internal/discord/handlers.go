package discord

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"strings"

	"github.com/disgoorg/disgo/bot"
	disgo "github.com/disgoorg/disgo/discord"
	"github.com/disgoorg/disgo/events"
	"github.com/mauriciofsnts/bot/internal/config"
	"github.com/mauriciofsnts/bot/internal/discord/ctx"
	customEvents "github.com/mauriciofsnts/bot/internal/discord/events"
)

func OnMessageCreatedEvent(event *events.MessageCreate, client bot.Client, cfg config.Config) {
	message := event.Message

	if message.Author.Bot {
		return
	}

	if !strings.HasPrefix(message.Content, cfg.Discord.Prefix) {
		return
	}

	inputMessage := strings.Split(message.Content, " ")

	commandName := strings.TrimPrefix(inputMessage[0], cfg.Discord.Prefix)
	found, cmd := ctx.GetCommandByAlias(commandName)

	if !found {
		// in this case is possible because is the user that is sending the message
		return
	}

	args := inputMessage[1:]

	trigger := ctx.TriggerEvent{
		AuthorId:       message.Author.ID,
		ChannelId:      message.ChannelID,
		MessageId:      message.ID,
		EventTimestamp: message.CreatedAt,
	}

	guildId := message.GuildID

	if guildId != nil {
		trigger.GuildId = *guildId
	}

	msg := ctx.Execute(args, cmd, trigger, ctx.MESSAGE, StartedAt, client)

	if msg != nil {
		msg.MessageReference = &disgo.MessageReference{MessageID: &message.ID}
		event.Client().Rest().CreateMessage(event.ChannelID, *msg)
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
	id := event.ComponentInteraction.Data.CustomID()
	found, component := ctx.GetComponentState(id)

	if !found {
		slog.Error("Button state not found: ", slog.String("id", id))
		return
	}

	component.Handler(event, &component.State)
}
