package discord

import (
	"log/slog"
	"strings"

	"github.com/disgoorg/disgo/discord"
	"github.com/disgoorg/disgo/events"
	"github.com/mauriciofsnts/exodia/internal/config"
	"github.com/mauriciofsnts/exodia/internal/discord/commands"
)

func OnMessageCreatedEvent(event *events.MessageCreate) {
	message := event.Message

	if message.Author.Bot {
		return
	}

	if !strings.HasPrefix(message.Content, config.Envs.Discord.Prefix) {
		return
	}

	msg := strings.Split(message.Content, " ")

	commandName := strings.TrimPrefix(msg[0], config.Envs.Discord.Prefix)
	found, _ := commands.SearchCommandByAlias(commandName)

	if found == false {
		// in this case is possible because is the user that is sending the message
		// return
	}

	args := msg[1:]

	slog.Info("Args: ", slog.String("args", strings.Join(args, " ")))

	if commandName == "ping" {
		event.Client().Rest().CreateMessage(event.ChannelID, discord.MessageCreate{
			Content: "Pong!",
		})
	}

}

func OnInteractionCreatedEvent(event *events.ApplicationCommandInteractionCreate) {
	data := event.SlashCommandInteractionData()

	commandName := data.CommandName()

	found, _ := commands.SearchCommandByAlias(commandName)

	if found == false {
		slog.Error("Command not found: ", slog.String("command", commandName))
		slog.Error("The user can't dispatch an unknown slash command. Check if the command is registered")
		// return
	}

	slog.Debug("Command found: ", slog.String("command", commandName))

	// trigger := commands.TriggerEvent{
	// 	GuildId:        event.GuildID().String(),
	// 	ChannelId:      event.Channel().ID().String(),
	// 	EventTimestamp: event.CreatedAt(),
	// }

	var commandArgs []string

	for _, option := range data.Options {
		commandArgs = append(commandArgs, string(option.Value))
	}

	//TODO! Implement the command execution
	if commandName == "ping" {
		event.CreateMessage(discord.NewMessageCreateBuilder().SetContent(data.String("pong")).Build())
	}

}

func OnReadyEvent(event *events.Ready) {
	slog.Info("Bot is ready!")
}
