package ctx

import (
	"github.com/disgoorg/disgo/bot"
	"github.com/disgoorg/disgo/events"
	"github.com/mauriciofsnts/bot/internal/database/models"
	"github.com/mauriciofsnts/bot/internal/providers"
)

type ComponentState struct {
	TriggerEvent TriggerEvent
	Client       bot.Client
	State        map[string]any
}

type ComponentHandler func(event *events.ComponentInteractionCreate, ctx *ComponentState)

type Component struct {
	State   ComponentState
	Handler ComponentHandler
}

func GetComponentStateInDatabase(messageId string) (bool, models.GuildState) {
	state, err := providers.Services.GuildState.GetComponentStateByMessageID(messageId)

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
