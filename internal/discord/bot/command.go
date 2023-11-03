package bot

import (
	"github.com/diamondburned/arikawa/v3/discord"
	"github.com/pauloo27/logger"
)

type BaseType struct {
	Name string
}

type ParameterType struct {
	Name     string
	BaseType *BaseType
}

type Parameter struct {
	ValidValues []any
	Name        string
	Type        *ParameterType
	Required    bool
}

type Handler func(*Context) discord.Embed

type Command struct {
	Parameters []*Parameter
	Aliases    []string
	Name       string
	Handler    Handler
}

var cmnd = make(map[string]Command)

func RegisterCommand(name string, command Command) {
	cmnd[name] = command
}

func GetCommand(name string) (*Command, bool) {
	command, found := cmnd[name]
	return &command, found
}

func getParsedOptions(command *Command) []discord.CommandOption {
	var options []discord.CommandOption

	for _, param := range command.Parameters {

		switch param.Type.BaseType {
		case TypeString:
			options = append(options, &discord.StringOption{
				OptionName:  param.Name,
				Description: "what's echoed back",
				Required:    param.Required,
			})
		case TypeInt:
			options = append(options, &discord.IntegerOption{
				OptionName:  param.Name,
				Description: "what's echoed back",
				Required:    param.Required,
			})
		case TypeBool:
			options = append(options, &discord.BooleanOption{
				OptionName:  param.Name,
				Description: "what's echoed back",
				Required:    param.Required,
			})
		default:
			logger.Fatal("Invalid type")
		}

	}

	return options
}
