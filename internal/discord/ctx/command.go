package ctx

import (
	"log/slog"

	"github.com/disgoorg/disgo/bot"
	"github.com/disgoorg/disgo/discord"
	"github.com/disgoorg/disgo/events"
	"github.com/mauriciofsnts/bot/internal/models"
	"github.com/mauriciofsnts/bot/internal/providers"
)

type Command struct {
	Name             string
	Description      string
	Aliases          []string
	Options          []discord.ApplicationCommandOption
	Handler          func(ctx Context) *discord.MessageCreate
	ComponentHandler func(event *events.ComponentInteractionCreate, ctx *ComponentState)
}

var commands = make(map[string]Command)

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

// Componente
type ComponentState struct {
	TriggerEvent TriggerEvent
	Client       bot.Client
	State        []any
}

type ComponentHandler func(event *events.ComponentInteractionCreate, ctx *ComponentState)

type Component struct {
	State   ComponentState
	Handler ComponentHandler
}

var buttonState = make(map[string]Component)

func RegisterComponent(id string, component Component) {
	buttonState[id] = component
}

func UpdateComponentState(id string, state []any) {
	if component, ok := buttonState[id]; ok {
		component.State.State = state
		buttonState[id] = component
	}
}

func GetComponentState(id string) (bool, Component) {
	component, ok := buttonState[id]
	return ok, component
}

func GetComponentStateInDatabase(id string) (bool, models.GuildState) {
	state, err := providers.Services.GuildState.GetComponentStateById(id)

	if err != nil {
		return false, models.GuildState{}
	}

	return true, *state
}

func GetComponentHandlerByName(id string) (bool, ComponentHandler) {
	if component, ok := commands[id]; ok {
		return true, component.ComponentHandler
	}

	return false, nil
}

func RemoveComponent(id string) {
	delete(buttonState, id)
}
