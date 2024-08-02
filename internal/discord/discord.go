package discord

import (
	"context"
	"log/slog"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/disgoorg/disgo"
	"github.com/disgoorg/disgo/bot"
	"github.com/disgoorg/disgo/discord"
	"github.com/disgoorg/disgo/events"
	"github.com/disgoorg/disgo/gateway"
	"github.com/mauriciofsnts/bot/internal/config"

	_ "github.com/mauriciofsnts/bot/internal/discord/commands"
	"github.com/mauriciofsnts/bot/internal/discord/ctx"
)

var StartedAt time.Time

func Init() {
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
				gateway.IntentGuildMessageReactions,
			),
			gateway.WithPresenceOpts(gateway.WithListeningActivity("your bullshit", gateway.WithActivityState("lol")), gateway.WithOnlineStatus(discord.OnlineStatusDND)),
		),
		bot.WithEventListenerFunc(OnReadyEvent),
	)

	if err != nil {
		slog.Error("Error initializing Discord client: ", slog.Any("error", err))
		panic(err)
	}

	client.AddEventListeners(&events.ListenerAdapter{
		OnMessageCreate: func(event *events.MessageCreate) {
			OnMessageCreatedEvent(event, &client)
		},
		OnApplicationCommandInteraction: func(event *events.ApplicationCommandInteractionCreate) {
			OnInteractionCreatedEvent(event, &client)
		},
		OnGuildChannelCreate: func(event *events.GuildChannelCreate) {
			OnGuildChannelCreatedEvent(event, &client)
		},
		OnMessageReactionAdd: func(event *events.MessageReactionAdd) {
			OnMessageReactionAddedEvent(event, &client)
		},
		OnComponentInteraction: func(event *events.ComponentInteractionCreate) {
			OnComponentInteractionEvent(event, &client)
		},
	})

	defer client.Close(context.Background())

	if config.Envs.Discord.SyncCommands {
		ctx.SyncCommands(client)
	}

	// connect to the gateway
	if err = client.OpenGateway(context.Background()); err != nil {
		slog.Error("Error while connecting to the gateway: ", slog.Any("error", err))
		panic(err)
	}

	StartedAt = time.Now()

	slog.Info("Bot is running. Press CTRL+C to exit.")
	s := make(chan os.Signal, 1)
	signal.Notify(s, syscall.SIGINT, syscall.SIGTERM)
	<-s
}
