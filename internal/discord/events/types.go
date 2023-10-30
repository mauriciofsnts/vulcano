package events

import (
	"github.com/bwmarrin/discordgo"
	"github.com/mauriciofsnts/vulcano/internal/i18n"
)

type Interaction struct {
	Interaction *discordgo.InteractionCreate
	Args        []*discordgo.ApplicationCommandInteractionDataOption
}

type Message struct {
	Message *discordgo.MessageCreate
	Args    []string
}

type CommandMessage struct {
	CommandHandler *CommandHandler
	T              *i18n.Language
	Session        *discordgo.Session
	Guild          *discordgo.Guild
	Message        *Message
	Interaction    *Interaction
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
