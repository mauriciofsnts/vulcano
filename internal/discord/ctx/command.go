package ctx

type Command struct {
	Name        string
	Description string
	Aliases     []string
	Handler     func(ctx *Context)
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
