package discord

import (
	"context"
	"log/slog"
	"os"
	"os/signal"
	"syscall"

	"github.com/disgoorg/disgo"
	"github.com/disgoorg/disgo/bot"
	"github.com/disgoorg/disgo/gateway"
	"github.com/mauriciofsnts/exodia/internal/config"

	_ "github.com/mauriciofsnts/exodia/internal/discord/commands"
	"github.com/mauriciofsnts/exodia/internal/discord/ctx"
)

func InitDiscord() {
	slog.Debug("Initializing Bot...")
	slog.Debug("Disgo version ", slog.String("version", disgo.Version))

	// Initialize Discord
	client, err := disgo.New(config.Envs.Discord.Token,
		bot.WithGatewayConfigOpts(
			gateway.WithIntents(
				gateway.IntentGuilds,
				gateway.IntentGuildMessages,
				gateway.IntentDirectMessages,
				gateway.IntentGuildVoiceStates,
				gateway.IntentMessageContent,
			),
		),
		bot.WithEventListenerFunc(OnMessageCreatedEvent),
		// bot.WithEventListenerFunc(OnInteractionCreatedEvent),
		bot.WithEventListenerFunc(OnReadyEvent),
	)

	if err != nil {
		slog.Error("Error initializing Discord client: ", slog.Any("error", err))
		panic(err)
	}

	defer client.Close(context.Background())

	if config.Envs.Discord.SyncCommands {
		ctx.SyncCommands(client)
	}

	// connect to the gateway
	if err = client.OpenGateway(context.Background()); err != nil {
		slog.Error("Error while connecting to the gateway: ", slog.Any("error", err))
		panic(err)
	}

	slog.Info("Bot is running. Press CTRL+C to exit.")
	s := make(chan os.Signal, 1)
	signal.Notify(s, syscall.SIGINT, syscall.SIGTERM)
	<-s
}
