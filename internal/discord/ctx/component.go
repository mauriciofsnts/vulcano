package ctx

import (
	"context"
	"encoding/json"
	"log/slog"
	"time"

	"dario.cat/mergo"

	"github.com/disgoorg/disgo/bot"
	"github.com/disgoorg/disgo/events"
	"github.com/mauriciofsnts/bot/internal/providers"
)

type DiscordComponentContext struct {
	TriggerEvent TriggerEvent
	Client       bot.Client
	State        map[string]any
}

type ComponentHandler func(event *events.ComponentInteractionCreate, ctx *DiscordComponentContext)

type Component struct {
	State   DiscordComponentContext
	Handler ComponentHandler
}

type GuildState struct {
	ComponentID    string         `json:"component_id"`
	GuildID        string         `json:"guild_id"`
	AuthorID       string         `json:"author_id"`
	ChannelID      string         `json:"channel_id"`
	MessageID      string         `json:"message_id"`
	EventTimestamp time.Time      `json:"event_timestamp"`
	Command        string         `json:"command"`
	State          map[string]any `json:"state" gorm:"serializer:json"`
	Ttl            time.Time      `json:"ttl"`
}

func GetComponentHandler(id string) (bool, ComponentHandler) {
	if component, ok := commands[id]; ok {
		return true, component.ComponentHandler
	}

	return false, nil
}

func CreateCommandState(id string, data GuildState, ttl time.Duration) bool {
	err := providers.Cache.Set(context.Background(), id, data, ttl)

	if err != nil {
		slog.Error("Error creating command state", "err", err)
		return true
	}

	return false
}

func UpdateCommandState(id string, newData GuildState) bool {
	state, found := GetCommandState(id)

	if !found {
		slog.Error("State not found")
		return true
	}

	new := &state
	err := mergo.Merge(new, newData, mergo.WithOverride)

	if err != nil {
		slog.Error("Error merging data", "err", err)
		return true
	}

	err = providers.Cache.Set(context.Background(), id, new, 1*time.Hour)

	if err != nil {
		slog.Error("Error creating command state", "err", err)
		return true
	}

	return false
}

func GetCommandState(id string) (GuildState, bool) {
	data, err := providers.Cache.Get(context.Background(), id)

	if err != nil {
		return GuildState{}, false
	}

	var response GuildState
	err = json.Unmarshal([]byte(data), &response)

	if err != nil {
		return GuildState{}, false
	}

	return response, true
}

func DeleteCommandState(id string) bool {
	err := providers.Cache.Delete(context.Background(), id)
	if err != nil {
		slog.Error("Error deleting valkey cache:", "error", err)
		return false
	}
	return true
}
