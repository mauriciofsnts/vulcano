package events

var cmd = make(map[string]CommandInfo)

func New(config Config) (ch *CommandHandler) {
	ch = &CommandHandler{
		commands: cmd,
		config:   config,
	}

	return
}

// Register registers a command to be handled by the command handler.
func Register(name string, commandInfo CommandInfo) {
	cmd[name] = commandInfo
}

// Get retrieves the Command (Data type) from the CommandMap map.
func (ch CommandHandler) Get(name string) (*CommandInfo, bool) {
	commandInfo, found := ch.commands[name]
	return &commandInfo, found
}
