package bot

import (
	"context"
	"time"

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
		Prefix: config.Vulcano.Prefix,
	}
}

type Bot struct {
	config    Configuration
	t         i18n.Language
	State     *state.State
	StartedAt time.Time
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

	bot.Use(cmdroute.Deferrable(bot.State, cmdroute.DeferOpts{}))
	bot.State.AddInteractionHandler(bot.Router)
	bot.State.AddIntents(gateway.IntentGuildMessages)
	bot.State.AddIntents(gateway.IntentDirectMessages)

	bot.InitHandler()

	if err := bot.State.Open(context.Background()); err != nil {
		logger.Debug("Failed to open state:", err)
	}

	return bot, nil
}

func (bot *Bot) Close() {
	bot.State.Close()
}
