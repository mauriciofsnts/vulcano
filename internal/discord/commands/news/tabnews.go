package news

import (
	"fmt"
	"log/slog"
	"strconv"
	"sync"
	"time"

	"github.com/disgoorg/disgo/discord"
	"github.com/disgoorg/disgo/events"
	"github.com/mauriciofsnts/bot/internal/database/models"
	"github.com/mauriciofsnts/bot/internal/discord/ctx"
	"github.com/mauriciofsnts/bot/internal/providers"
	"github.com/mauriciofsnts/bot/internal/providers/news"
	"github.com/mauriciofsnts/bot/internal/utils"
)

func init() {
	ctx.RegisterCommand("tabnews", ctx.Command{
		Name:        "tabnews",
		Aliases:     []string{"tn", "tabnews"},
		Description: ctx.Translate().Commands.Tabnews.Description.Str(),
		Options: []discord.ApplicationCommandOption{
			discord.ApplicationCommandOptionInt{
				Name:        "page",
				Description: "The page number you want to see",
				Required:    false,
				MinValue:    utils.PtrTo(1),
				MaxValue:    utils.PtrTo(99),
			},
		},
		Handler: func(data ctx.CommandExecutionContext) *discord.MessageCreate {
			page := 1

			if len(data.Args) > 0 {
				parsedPage, err := strconv.Atoi(data.Args[0])
				if err != nil {
					page = 1
				}
				page = parsedPage
			}

			fields, err := fetchNews(page)

			if err != nil {
				reply := data.Response.BuildDefaultErrorMessage(err)
				return &reply
			}

			messageBuilder := discord.NewMessageCreateBuilder()
			embedBuilder := discord.NewEmbedBuilder().
				SetTitle("Tabnews").
				SetDescription(ctx.Translate().Commands.Tabnews.Reply.Str()).
				SetColor(0xffffff).
				SetFields(fields...).
				SetFooter(fmt.Sprintf("Page %d", page), "")
			embed := embedBuilder.Build()
			messageBuilder.SetEmbeds(embed)

			actionButtonId := fmt.Sprintf("tabnews-next-%d", data.TriggerEvent.MessageId)
			prevButtonId := fmt.Sprintf("tabnews-prev-%d", data.TriggerEvent.MessageId)

			if page > 1 {
				messageBuilder.AddActionRow(
					discord.NewSecondaryButton("⬅️", prevButtonId),
					discord.NewSecondaryButton("➡️", actionButtonId),
				)
			} else {
				messageBuilder.AddActionRow(discord.NewSecondaryButton("➡️", actionButtonId))
			}

			msg := messageBuilder.Build()

			createdMessage, err := data.Client.Rest().CreateMessage(data.TriggerEvent.ChannelId, msg)

			if err != nil {
				slog.Error("Error creating message", "err", err.Error())
				return nil
			}

			providers.Services.GuildState.CreateComponentState(&models.GuildState{
				GuildID:        data.TriggerEvent.GuildId.String(),
				ComponentID:    actionButtonId,
				AuthorID:       data.TriggerEvent.AuthorId.String(),
				ChannelID:      data.TriggerEvent.ChannelId.String(),
				MessageID:      createdMessage.ID.String(),
				Command:        "tabnews",
				State:          map[string]any{"page": page},
				Ttl:            time.Now(),
				EventTimestamp: data.TriggerEvent.EventTimestamp,
			})

			return nil
		},
		ComponentHandler: func(event *events.ComponentInteractionCreate, data *ctx.DiscordComponentContext) {
			actionButtonId := fmt.Sprintf("tabnews-next-%d", data.TriggerEvent.MessageId)
			prevButtonId := fmt.Sprintf("tabnews-prev-%d", data.TriggerEvent.MessageId)

			page := int(data.State["page"].(float64))
			nextPageState := page

			if event.ComponentInteraction.Data.CustomID() == prevButtonId {
				nextPageState = max(page-1, 1)
			} else {
				nextPageState = page + 1
			}

			fields, _ := fetchNews(nextPageState)

			embedBuilder := discord.NewEmbedBuilder().
				SetTitle("Tabnews").
				SetDescription(ctx.Translate().Commands.Tabnews.Reply.Str()).
				SetColor(0xffffff).
				SetFooter(fmt.Sprintf("Page %d", nextPageState), "").
				SetFields(fields...)
			embed := embedBuilder.Build()
			embeds := []discord.Embed{embed}

			providers.Services.GuildState.UpdateComponentState(data.TriggerEvent.MessageId.String(), map[string]any{"page": nextPageState})

			newActionRow := discord.NewActionRow()
			newActionRow.AddComponents(
				discord.NewSecondaryButton("⬅️", prevButtonId),
				discord.NewSecondaryButton("➡️", actionButtonId),
			)

			components := discord.ActionRowComponent{
				discord.NewSecondaryButton("➡️", actionButtonId),
			}

			if nextPageState > 1 {
				components = discord.ActionRowComponent{
					discord.NewSecondaryButton("⬅️", prevButtonId),
					discord.NewSecondaryButton("➡️", actionButtonId),
				}
			}

			event.UpdateMessage(discord.MessageUpdate{
				Embeds:     &embeds,
				Components: &[]discord.ContainerComponent{components},
			})
		},
	})
}

func fetchNews(page int) ([]discord.EmbedField, error) {
	minPage := 1

	if page < minPage {
		page = minPage
	}

	tnArticles, err := providers.News.Tabnews(page, 12)

	if err != nil {
		return nil, err
	}

	fields := make([]discord.EmbedField, len(tnArticles))
	var wg sync.WaitGroup

	for i, article := range tnArticles {
		wg.Add(1)
		go func(idx int, article news.TabnewsArticle) {
			defer wg.Done()

			url, _ := providers.Shorten.ShortenLink(fmt.Sprintf("https://www.tabnews.com.br/%s/%s", article.Owner_username, article.Slug), nil)

			value := fmt.Sprintf("⭐ %d · %s · %s", article.Tabcoins, article.Owner_username, url)
			fields[idx] = discord.EmbedField{
				Name:  article.Title,
				Value: value,
			}
		}(i, article)
	}

	wg.Wait()
	return fields, nil
}
