package bot

import (
	"github.com/diamondburned/arikawa/v3/discord"
)

type Handler func(*Context) discord.Embed

type Command struct {
	Name        string
	Description string
	Aliases     []string
	Parameters  []discord.CommandOption
	Handler     Handler
}

var cmnd = make(map[string]Command)

func RegisterCommand(name string, command Command) {
	cmnd[name] = command
}

func GetCommand(name string) (*Command, bool) {
	command, found := cmnd[name]
	return &command, found
}

func GetCommandByNameAndAliases(name string) (*Command, bool) {
	for _, command := range cmnd {
		if command.Name == name {
			return &command, true
		}
		for _, alias := range command.Aliases {
			if alias == name {
				return &command, true
			}
		}
	}
	return nil, false
}
