package bot

import (
	"github.com/bwmarrin/discordgo"
	"github.com/mauriciofsnts/vulcano/internal/config"
	"github.com/mauriciofsnts/vulcano/internal/discord/events"
	"github.com/pauloo27/logger"
)

type Bot struct {
	config         Configuration
	session        *discordgo.Session
	commandHandler *events.CommandHandler
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
	}

	bot.session, err = discordgo.New("Bot " + config.Token)

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

	return
}

// Close closes the bot
func (bot Bot) Close() {
	bot.session.Close()
}
