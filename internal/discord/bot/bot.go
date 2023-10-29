package bot

import (
	"github.com/bwmarrin/discordgo"
	"github.com/mauriciofsnts/vulcano/internal/config"
	"github.com/mauriciofsnts/vulcano/internal/discord/events"
	"github.com/mauriciofsnts/vulcano/internal/i18n"
	"github.com/pauloo27/logger"
)

type Bot struct {
	config         Configuration
	session        *discordgo.Session
	commandHandler *events.CommandHandler
	t              *i18n.Language
}

type Configuration struct {
	Token  string `required:"true"`
	Prefix string `default:"!"`
}

func New() (bot *Bot, err error) {

	config := Configuration{
		Token:  config.Vulcano.Token,
		Prefix: ",",
	}

	bot = &Bot{
		config: config,
		t:      i18n.GetLanguage("pt_BR"),
	}

	bot.session, err = discordgo.New("Bot " + config.Token)
	bot.session.Identify.Intents = discordgo.IntentsAllWithoutPrivileged | discordgo.IntentsGuildMembers

	if err != nil {
		return nil, err
	}

	err = bot.session.Open()

	if err != nil {
		return nil, err
	}

	logger.Info("Bot is now running.  Press CTRL-C to exit.")

	bot.commandHandler = events.New(events.Config{
		Prefix: config.Prefix,
	})

	bot.InitHandler()

	applicationCommands := make([]*discordgo.ApplicationCommand, len(bot.commandHandler.Commands))

	i := 0
	for _, command := range bot.commandHandler.Commands {
		applicationCommands[i] = command.ApplicationCommand
		i++
	}

	_, err = bot.session.ApplicationCommandBulkOverwrite(bot.session.State.User.ID, "", applicationCommands)

	if err != nil {
		return nil, err
	}

	return
}

// Close closes the bot
func (bot Bot) Close() {
	bot.session.Close()
}
