package ctx

import "github.com/mauriciofsnts/bot/internal/i18n"

func Translate() i18n.Language {
	return *i18n.GetLanguage("fenix")
}
