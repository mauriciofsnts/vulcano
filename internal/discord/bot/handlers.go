package bot

import (
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/mauriciofsnts/vulcano/internal/discord/events"
)

func (bot Bot) onMessageCreate(session *discordgo.Session, message *discordgo.MessageCreate) {
	if message.Author.ID == session.State.User.ID {
		return
	}

	if strings.HasPrefix(message.Content, bot.config.Prefix) {
		arguments := strings.Fields(strings.TrimPrefix(message.Content, bot.config.Prefix))
		cmdName := arguments[0]

		arguments = arguments[1:]

		channel, err := session.Channel(message.ChannelID)
		if err != nil {
			return
		}

		guild, err := session.Guild(channel.GuildID)

		if err != nil {
			return
		}

		commandMessage := events.CommandMessage{
			CommandHandler: bot.commandHandler,
			T:              bot.t,
			Session:        session,
			Guild:          guild,
			Message:        &events.M{Message: message, Args: arguments},
			Interaction:    nil,
		}

		commandInfo, found := bot.commandHandler.Get(cmdName)
		if !found {
			return
		}

		cmdFunction := commandInfo.Function

		cmdFunction(commandMessage)
	}
}

func (bot Bot) onInteractionCreate(session *discordgo.Session, interaction *discordgo.InteractionCreate) {
	if interaction.Member.User.Bot {
		return
	}

	channel, err := session.Channel(interaction.ChannelID)

	if err != nil {
		return
	}

	guild, err := session.Guild(channel.GuildID)
	if err != nil {
		return
	}

	cmd, found := bot.commandHandler.Get(interaction.ApplicationCommandData().Name)

	if !found {
		return
	}

	commandInteraction := events.CommandMessage{
		CommandHandler: bot.commandHandler,
		T:              bot.t,
		Session:        session,
		Guild:          guild,
		Message:        nil,
		Interaction:    &events.I{Interaction: interaction, Args: interaction.ApplicationCommandData().Options},
	}

	cmd.Function(commandInteraction)

}

func (bot Bot) InitHandler() {
	bot.session.AddHandler(bot.onMessageCreate)
	bot.session.AddHandler(bot.onInteractionCreate)
}
