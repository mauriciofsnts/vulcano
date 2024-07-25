package ctx

import (
	"time"
)

type EventType string

const (
	SLASH_COMMAND EventType = "SLASH_COMMAND"
	MESSAGE       EventType = "MESSAGE"
)

type TriggerEvent struct {
	AuthorId       string
	ChannelId      string
	GuildId        string
	MessageId      string
	EventTimestamp time.Time
}

type Context struct {
	StartTimestamp time.Time
	TriggerEvent   TriggerEvent
	Type           EventType
	Args           []string
}

func Execute(
	args []string,
	command *Command,
	trigger TriggerEvent,
	eventType EventType,
) {
	ctx := &Context{
		StartTimestamp: time.Now(),
		TriggerEvent:   trigger,
		Type:           eventType,
		Args:           args,
	}

	command.Handler(ctx)
}
