package events

import (
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/pauloo27/logger"
)

var cmd = make(map[string]CommandInfo)

func New(config Config) (ch *CommandHandler) {
	// list of commands
	logger.Info("Registering commands...", cmd)

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

func (ch CommandHandler) Process(session *discordgo.Session, message *discordgo.MessageCreate) {
	if message.Author.ID == session.State.User.ID {
		return
	}

	if strings.HasPrefix(message.Content, ch.config.Prefix) {
		arguments := strings.Fields(strings.TrimPrefix(message.Content, ch.config.Prefix))
		cmdName := arguments[0]

		// Removing the command from the arguments slice
		arguments = arguments[1:]

		// Get the guild in which the message was sent
		channel, err := session.Channel(message.ChannelID)
		if err != nil {
			return
		}

		guild, err := session.Guild(channel.GuildID)

		if err != nil {
			return
		}

		commandMessage := CommandMessage{
			CommandHandler: &ch,
			Session:        session,
			Guild:          guild,
			Message:        message,
			Args:           arguments,
			Interaction:    nil,
		}

		commandInfo, found := ch.Get(cmdName)
		if !found {
			return
		}

		cmdFunction := commandInfo.Function

		// Call the command's function
		cmdFunction(commandMessage)
	}
}
