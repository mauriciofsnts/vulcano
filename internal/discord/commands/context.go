package commands

import "time"

type TriggerEvent struct {
	AuthorId       string
	ChannelId      string
	GuildId        string
	EventTimestamp time.Time
}

type Context struct {
	StartTimestamp time.Time
	TriggerEvent   TriggerEvent
	Args           []string
}
