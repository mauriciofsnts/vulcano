package ctx

import "github.com/disgoorg/disgo/discord"

func Reply(
	title string,
	description string,
	fields []discord.EmbedField,
) discord.MessageCreate {
	builder := discord.NewMessageCreateBuilder()

	embedBuilder := discord.NewEmbedBuilder()
	embedBuilder.
		SetTitle(title).
		SetDescription(description).
		SetColor(0xffffff).
		SetFields(fields...)
	embed := embedBuilder.Build()

	builder.SetEmbeds(embed)
	return builder.Build()
}

func ReplyErr(
	err error,
) discord.MessageCreate {
	builder := discord.NewMessageCreateBuilder()

	embedBuilder := discord.NewEmbedBuilder()
	embedBuilder.
		SetTitle("An error occurred").
		SetDescription(err.Error()).
		SetColor(0xff0000)

	embed := embedBuilder.Build()
	builder.SetEmbeds(embed)
	return builder.Build()
}
