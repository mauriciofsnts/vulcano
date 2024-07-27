package ctx

import (
	"log/slog"

	"github.com/disgoorg/disgo/bot"
	"github.com/disgoorg/disgo/discord"
)

type Command struct {
	Name        string
	Description string
	Aliases     []string
	Options     []discord.ApplicationCommandOption
	Handler     func(ctx *Context) discord.MessageCreate
}

var commands = make(map[string]Command)

func AttachCommand(name string, cmd Command) {
	commands[name] = cmd
}

func SearchCommandByAlias(alias string) (bool, *Command) {
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

func ParseCommandsToSlashCommands() []discord.ApplicationCommandCreate {
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

func SyncCommands(client bot.Client) {
	commands := ParseCommandsToSlashCommands()
	var err error

	if _, err = client.Rest().SetGlobalCommands(client.ApplicationID(), commands); err != nil {
		slog.Error("error while registering commands: ", err)
	}
}
