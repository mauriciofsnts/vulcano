package slash

import (
	"github.com/bwmarrin/discordgo"
	"github.com/mauriciofsnts/vulcano/internal/i18n"
	"github.com/pauloo27/logger"
)

type DiscordContext struct {
	S *discordgo.Session
	I *discordgo.InteractionCreate
}

type SlashCommand struct {
	*discordgo.ApplicationCommand
	Handler func(ctx *DiscordContext, t *i18n.Language)
}

var commands = make(map[string]*SlashCommand)

func RegisterSlashCommand(cmds ...*SlashCommand) {
	for _, command := range cmds {
		commands[command.Name] = command
	}
}

func Start(s *discordgo.Session) error {

	applicationCommands := make([]*discordgo.ApplicationCommand, len(commands))

	i := 0
	for _, c := range commands {
		logger.Debugf("Registering command: %s", c.Name)

		applicationCommands[i] = c.ApplicationCommand
		i++
	}

	s.AddHandler(func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		var lang i18n.EnumLanguage

		commandName := i.ApplicationCommandData().Name

		if command, ok := commands[commandName]; ok {
			command.Handler(
				&DiscordContext{
					S: s,
					I: i,
				},
				i18n.GetLanguage(lang),
			)
		}

	})

	_, err := s.ApplicationCommandBulkOverwrite(s.State.User.ID, "", applicationCommands)

	return err
}
