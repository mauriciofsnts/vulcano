package bot

import (
	"github.com/bwmarrin/discordgo"
	"github.com/mauriciofsnts/vulcano/internal/discord/events"
)

func (bot Bot) onMessageCreate(session *discordgo.Session, info *discordgo.MessageCreate) {
	if !info.Author.Bot {
		bot.commandHandler.Process(session, info)
	}
}

func (bot Bot) onInteractionCreate(session *discordgo.Session, info *discordgo.InteractionCreate) {
	if info.Member.User.Bot {
		return
	}

	channel, err := session.Channel(info.ChannelID)

	if err != nil {
		return
	}

	guild, err := session.Guild(channel.GuildID)
	if err != nil {
		return
	}

	cmd, found := bot.commandHandler.Get(info.ApplicationCommandData().Name)

	if !found {
		return
	}

	commandInteraction := events.CommandMessage{
		CommandHandler: bot.commandHandler,
		Session:        session,
		Guild:          guild,
		Message:        nil,
		Interaction:    &events.I{Interaction: info, Args: info.ApplicationCommandData().Options},
	}

	cmd.Function(commandInteraction)

}

func (bot Bot) InitHandler() {
	bot.session.AddHandler(bot.onMessageCreate)
	bot.session.AddHandler(bot.onInteractionCreate)
}
