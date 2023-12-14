package t

import "github.com/mauriciofsnts/vulcano/internal/i18n"

func Translate() i18n.Language {
	return *i18n.GetLanguage("pt_BR")
}
