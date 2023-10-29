package events

import (
	"github.com/bwmarrin/discordgo"
	"github.com/mauriciofsnts/vulcano/internal/i18n"
)

type I struct {
	Interaction *discordgo.InteractionCreate
	Args        []*discordgo.ApplicationCommandInteractionDataOption
}

type M struct {
	Message *discordgo.MessageCreate
	Args    []string
}

type CommandMessage struct {
	CommandHandler *CommandHandler
	T              *i18n.Language
	Session        *discordgo.Session
	Guild          *discordgo.Guild
	Message        *M
	Interaction    *I
}

type CommandInfo struct {
	*discordgo.ApplicationCommand
	Function Command
}

type Command func(CommandMessage)

type CommandMap map[string]CommandInfo

type CommandHandler struct {
	Commands CommandMap
	config   Config
}

type Config struct {
	Prefix string
}
