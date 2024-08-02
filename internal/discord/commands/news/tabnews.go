package news

import (
	"fmt"
	"log/slog"
	"strconv"
	"sync"

	"github.com/disgoorg/disgo/discord"
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

	actionButtonId := fmt.Sprintf("tabnews-%d", cmd.TriggerEvent.MessageId)
	messageBuilder.AddActionRow(discord.NewPrimaryButton("Next page", actionButtonId))
	msg := messageBuilder.Build()

	cmd.Client.Rest().CreateMessage(cmd.TriggerEvent.ChannelId, msg)

	ctx.AttachComponentState(actionButtonId, createComponentState(actionButtonId, page+1, cmd))

	return nil
}

func createComponentState(actionButtonId string, nextPage int, cmd *ctx.Context) ctx.Component {
	return ctx.Component{
		State: ctx.ComponentState{
			TriggerEvent: cmd.TriggerEvent,
			Client:       cmd.Client,
			State:        []interface{}{nextPage},
		},
		Handler: func(state *ctx.ComponentState) *[]discord.Embed {
			page := state.State[0].(int)
			fields, _ := fetchNews(page)
			embedBuilder := discord.NewEmbedBuilder().
				SetTitle("Latest news from Tabnews").
				SetDescription("Here are the latest news from the tabnews website").
				SetColor(0xffffff).
				SetFields(fields...)
			embed := embedBuilder.Build()
			embeds := []discord.Embed{embed}

			ctx.UpdateComponentStateById(actionButtonId, []interface{}{page + 1})
			return &embeds
		},
	}
}

func fetchNews(page int) ([]discord.EmbedField, error) {
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
				slog.Error("Error shortening url: ", err.Error())
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
