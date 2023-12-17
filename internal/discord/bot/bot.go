package bot

import (
	"context"
	"time"

	"github.com/diamondburned/arikawa/v3/api/cmdroute"
	"github.com/diamondburned/arikawa/v3/gateway"
	"github.com/diamondburned/arikawa/v3/state"
	"github.com/mauriciofsnts/vulcano/internal/config"
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

type Discord struct {
	config    Configuration
	State     *state.State
	StartedAt time.Time
	*cmdroute.Router
}

func New() (bot *Discord, err error) {
	instance := state.New("Bot " + config.Vulcano.Token)
	router := cmdroute.NewRouter()

	bot = &Discord{
		config: *NewConfiguration(),
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

func (bot *Discord) Close() error {
	return bot.State.Close()
}

func (bot *Discord) IsAlive() bool {
	me, _ := bot.State.Me()
	return me != nil
}
