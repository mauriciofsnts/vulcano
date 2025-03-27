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
	Handler          func(context CommandExecutionContext) *discord.MessageCreate
	ComponentHandler func(event *events.ComponentInteractionCreate, ctx *DiscordComponentContext)
}

func RegisterCommand(name string, cmd Command) {
	commands[name] = cmd
}

func GetCommandByNameOrAlias(alias string) (bool, *Command, string) {
	for _, command := range commands {
		if command.Name == alias {
			return true, &command, ""
		}
		for _, cmdalias := range command.Aliases {
			if cmdalias == alias {
				return true, &command, alias
			}
		}
	}
	return false, nil, ""
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
