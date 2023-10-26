package events

import "github.com/bwmarrin/discordgo"

type CommandMessage struct {
	CommandHandler *CommandHandler
	Session        *discordgo.Session
	Guild          *discordgo.Guild
	Message        *discordgo.MessageCreate
	Interaction    *discordgo.InteractionCreate
	Args           []string
}

type CommandInfo struct {
	*discordgo.ApplicationCommand
	Function Command
}

type Command func(CommandMessage)

type CommandMap map[string]CommandInfo

type CommandHandler struct {
	commands CommandMap
	config   Config
}

type Config struct {
	Prefix string
}
