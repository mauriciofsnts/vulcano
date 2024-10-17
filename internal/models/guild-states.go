package models

import (
	"time"

	"gorm.io/gorm"
)

type GuildState struct {
	gorm.Model
	ComponentID    string         `json:"component_id"`
	GuildID        string         `json:"guild_id"`
	AuthorID       string         `json:"author_id"`
	ChannelID      string         `json:"channel_id"`
	MessageID      string         `json:"message_id"`
	EventTimestamp time.Time      `json:"event_timestamp"`
	Command        string         `json:"command"` // Command name, e.g. "ping"
	State          map[string]any `json:"state" gorm:"serializer:json"`
	Ttl            time.Time      `json:"ttl"`
}
