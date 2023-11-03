package bot

import (
	"context"

	"github.com/diamondburned/arikawa/v3/api"
	"github.com/diamondburned/arikawa/v3/api/cmdroute"
	"github.com/diamondburned/arikawa/v3/gateway"
	"github.com/diamondburned/arikawa/v3/state"
	"github.com/mauriciofsnts/vulcano/internal/config"
	"github.com/mauriciofsnts/vulcano/internal/i18n"
	"github.com/pauloo27/logger"
)

type Configuration struct {
	Token  string `required:"true"`
	Prefix string `default:"!"`
}

func NewConfiguration() *Configuration {
	return &Configuration{
		Token:  config.Vulcano.Token,
		Prefix: ",",
	}
}

type Bot struct {
	config Configuration
	t      i18n.Language
	State  *state.State
	*cmdroute.Router
}

func New() (bot *Bot, err error) {

	instance := state.New("Bot " + config.Vulcano.Token)
	router := cmdroute.NewRouter()

	bot = &Bot{
		config: *NewConfiguration(),
		t:      *i18n.GetLanguage("pt_BR"),
		State:  instance,
		Router: router,
	}

	bot.State.AddInteractionHandler(bot.Router)

	bot.State.AddIntents(gateway.IntentGuildMessages)
	bot.State.AddIntents(gateway.IntentDirectMessages)

	if err := bot.State.Open(context.Background()); err != nil {
		logger.Debug("Failed to open state:", err)
	}

	// Automatically defer handles if they're slow.
	bot.Use(cmdroute.Deferrable(bot.State, cmdroute.DeferOpts{}))

	// Register events
	bot.InitHandler()

	// Register commands
	var discCommands []api.CreateCommandData

	for _, command := range cmnd {
		options := getParsedOptions(&command)

		discCommands = append(discCommands, api.CreateCommandData{
			Name:    command.Name,
			Options: options,
		})
	}

	// List a size of commands
	logger.Info("Registered", len(cmnd), "commands")
	logger.Info("Registered", len(discCommands), "commands")

	bot.State.BulkOverwriteCommands(bot.State.Ready().Application.ID, discCommands)

	return bot, nil
}

func (bot *Bot) Close() {
	bot.State.Close()
}
