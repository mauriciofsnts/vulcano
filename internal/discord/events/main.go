package events

import "github.com/pauloo27/logger"

var Cmd = make(map[string]CommandInfo)

func New(config Config) (ch *CommandHandler) {
	ch = &CommandHandler{
		Commands: Cmd,
		config:   config,
	}

	return
}

// Register registers a command to be handled by the command handler.
func Register(name string, commandInfo CommandInfo) {
	logger.Debug("Registering command: " + name)
	Cmd[name] = commandInfo
}

// Get retrieves the Command (Data type) from the CommandMap map.
func (ch CommandHandler) Get(name string) (*CommandInfo, bool) {
	commandInfo, found := ch.Commands[name]
	return &commandInfo, found
}
