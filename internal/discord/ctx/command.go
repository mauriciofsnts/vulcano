package ctx

import (
	"log/slog"

	"github.com/disgoorg/disgo/bot"
	"github.com/disgoorg/disgo/discord"
	"github.com/disgoorg/disgo/events"
)

var commands = make(map[string]Command)

type Command struct {
	Name             string
	Description      string
	Aliases          []string
	Options          []discord.ApplicationCommandOption
	Handler          func(ctx Context) *discord.MessageCreate
	ComponentHandler func(event *events.ComponentInteractionCreate, ctx *ComponentState)
}

func RegisterCommand(name string, cmd Command) {
	commands[name] = cmd
}

func GetCommandByAlias(alias string) (bool, *Command) {
	for _, command := range commands {
		if command.Name == alias {
			return true, &command
		}
		for _, cmdalias := range command.Aliases {
			if cmdalias == alias {
				return true, &command
			}
		}
	}
	return false, nil
}

func ConvertToSlashCommands() []discord.ApplicationCommandCreate {
	var slashCommands []discord.ApplicationCommandCreate

	for _, command := range commands {
		slashCommands = append(slashCommands, discord.SlashCommandCreate{
			Name:        command.Name,
			Description: command.Description,
			Options:     command.Options,
		})
	}

	return slashCommands
}

func SyncSlashCommands(client bot.Client) {
	slashCommands := ConvertToSlashCommands()

	if _, err := client.Rest().SetGlobalCommands(client.ApplicationID(), slashCommands); err != nil {
		slog.Info("Failed to register commands", slog.Any("error", err))
	}
}
