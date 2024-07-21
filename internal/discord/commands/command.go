package commands

type Command struct {
	Name        string
	Description string
	Aliases     []string
	Handler     func(ctx *Context)
}

var commands map[string]Command

func AttachCommand(name string, cmd Command) {
	commands[name] = cmd
}

func SearchCommandByAlias(alias string) (bool, *Command) {
	for _, cmd := range commands {
		for _, a := range cmd.Aliases {
			if a == alias {
				return true, &cmd
			}
		}
	}
	return false, nil
}
