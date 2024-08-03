package news

import (
	"fmt"
	"log/slog"
	"strconv"
	"sync"

	"github.com/disgoorg/disgo/discord"
	"github.com/disgoorg/disgo/events"
	"github.com/mauriciofsnts/bot/internal/discord/ctx"
	"github.com/mauriciofsnts/bot/internal/providers/news"
	"github.com/mauriciofsnts/bot/internal/providers/shorten"
	"github.com/mauriciofsnts/bot/internal/providers/utils"
)

func init() {
	ctx.AttachCommand("tabnews", ctx.Command{
		Name:        "Tabnews",
		Aliases:     []string{"tn", "tabnews"},
		Description: "Get the latest news from the tabnews website",
		Options: []discord.ApplicationCommandOption{
			discord.ApplicationCommandOptionInt{
				Name:        "page",
				Description: "The page number you want to see",
				Required:    false,
				MinValue:    utils.PtrTo(1),
				MaxValue:    utils.PtrTo(99),
			},
		},
		Handler: handleTabnewsCommand,
	})
}

func handleTabnewsCommand(cmd *ctx.Context) *discord.MessageCreate {
	page := 1
	if len(cmd.Args) > 0 {
		if value, err := strconv.Atoi(cmd.Args[0]); err == nil && value >= 1 {
			page = value
		}
	}

	fields, err := fetchNews(page)
	if err != nil {
		reply := cmd.Response.ReplyErr(err)
		return &reply
	}

	messageBuilder := discord.NewMessageCreateBuilder()
	embedBuilder := discord.NewEmbedBuilder().
		SetTitle("Latest news from Tabnews").
		SetDescription("Here are the latest news from the tabnews website").
		SetColor(0xffffff).
		SetFields(fields...)
	embed := embedBuilder.Build()
	messageBuilder.SetEmbeds(embed)

	actionButtonId := fmt.Sprintf("tabnews-next-%d", cmd.TriggerEvent.MessageId)
	prevButtonId := fmt.Sprintf("tabnews-prev-%d", cmd.TriggerEvent.MessageId)

	messageBuilder.AddActionRow(discord.NewSecondaryButton("➡️", actionButtonId))

	msg := messageBuilder.Build()

	cmd.Client.Rest().CreateMessage(cmd.TriggerEvent.ChannelId, msg)

	ctx.AttachComponentState(actionButtonId, createComponentState(page+1, cmd))
	ctx.AttachComponentState(prevButtonId, createComponentState(page-1, cmd))

	return nil
}

func createComponentState(nextPage int, cmd *ctx.Context) ctx.Component {
	return ctx.Component{State: ctx.ComponentState{
		TriggerEvent: cmd.TriggerEvent,
		Client:       cmd.Client,
		State:        []interface{}{nextPage},
	}, Handler: func(event *events.ComponentInteractionCreate, state *ctx.ComponentState) {
		page := state.State[0].(int)
		fields, _ := fetchNews(page)

		embedBuilder := discord.NewEmbedBuilder().
			SetTitle("Latest news from Tabnews").
			SetDescription("Here are the latest news from the tabnews website").
			SetColor(0xffffff).
			SetFooter(fmt.Sprintf("Page %d", page), "").
			SetFields(fields...)
		embed := embedBuilder.Build()
		embeds := []discord.Embed{embed}

		actionButtonId := fmt.Sprintf("tabnews-next-%d", cmd.TriggerEvent.MessageId)
		prevButtonId := fmt.Sprintf("tabnews-prev-%d", cmd.TriggerEvent.MessageId)

		ctx.UpdateComponentStateById(actionButtonId, []interface{}{page + 1})
		ctx.UpdateComponentStateById(prevButtonId, []interface{}{page - 1})

		newActionRow := discord.NewActionRow()
		newActionRow.AddComponents(discord.NewSecondaryButton("⬅️", prevButtonId), discord.NewSecondaryButton("➡️", actionButtonId))

		components := discord.ActionRowComponent{
			discord.NewSecondaryButton("➡️", actionButtonId),
		}

		if page > 1 {
			components = discord.ActionRowComponent{
				discord.NewSecondaryButton("⬅️", prevButtonId),
				discord.NewSecondaryButton("➡️", actionButtonId),
			}
		}

		event.UpdateMessage(discord.MessageUpdate{
			Embeds:     &embeds,
			Components: &[]discord.ContainerComponent{components},
		})

	}}
}

func fetchNews(page int) ([]discord.EmbedField, error) {
	minPage := 1

	if page < minPage {
		page = minPage
	}

	tnArticles, err := news.GetTnNews(page, 15)

	if err != nil {
		return nil, err
	}

	fields := make([]discord.EmbedField, len(tnArticles))
	var wg sync.WaitGroup

	for i, article := range tnArticles {
		wg.Add(1)
		go func(idx int, article news.TnArticle) {
			defer wg.Done()
			shortenedUrl, err := shorten.Shortner(fmt.Sprintf("https://www.tabnews.com.br/%s/%s", article.Owner_username, article.Slug), nil)
			if err != nil {
				slog.Error("error shortening url: ", "err", err.Error())
				return
			}

			value := fmt.Sprintf("⭐ %d · %s · %s", article.Tabcoins, article.Owner_username, shortenedUrl)
			fields[idx] = discord.EmbedField{
				Name:  article.Title,
				Value: value,
			}
		}(i, article)
	}

	wg.Wait()
	return fields, nil
}
