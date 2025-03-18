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

	eventHandler "github.com/mauriciofsnts/bot/internal/discord/events"
)

var StartedAt time.Time

func Init(cfg config.Config) {
	slog.Debug("Initializing Bot...")
	slog.Debug("Disgo version ", slog.String("version", disgo.Version))
	slog.Debug("Commands prefix: ", slog.String("prefix", cfg.Discord.Prefix))

	// Initialize Discord
	client, err := disgo.New(cfg.Discord.Token,
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
		bot.WithEventListenerFunc(eventHandler.OnReadyEvent),
	)

	if err != nil {
		slog.Error("Error initializing Discord client: ", slog.Any("error", err))
		panic(err)
	}

	client.AddEventListeners(&events.ListenerAdapter{
		OnMessageCreate: func(event *events.MessageCreate) {
			eventHandler.OnMessageCreatedEvent(event, client, cfg, StartedAt)
		},
		OnApplicationCommandInteraction: func(event *events.ApplicationCommandInteractionCreate) {
			eventHandler.OnInteractionCreatedEvent(event, client, cfg, StartedAt)
		},
		OnGuildChannelCreate: func(event *events.GuildChannelCreate) {
			eventHandler.OnGuildChannelCreatedEvent(event, client)
		},
		OnMessageReactionAdd: func(event *events.MessageReactionAdd) {
			eventHandler.OnMessageReactionAddedEvent(event, client)
		},
		OnComponentInteraction: func(event *events.ComponentInteractionCreate) {
			eventHandler.OnComponentInteractionEvent(event, client)
		},
		OnGuildVoiceJoin: func(event *events.GuildVoiceJoin) {
		},
		OnGuildReady: func(event *events.GuildReady) {
			eventHandler.OnGuildReady(event)
		},
	})

	defer client.Close(context.Background())

	if cfg.Discord.SyncCommands {
		ctx.SyncSlashCommands(client)
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
