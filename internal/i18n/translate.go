package i18n

import (
	"fmt"
	"os"
	"strings"

	"github.com/ghodss/yaml"
	"github.com/pauloo27/logger"
)

type EnumLanguage string

var (
	Languages   []EnumLanguage
	languageMap map[EnumLanguage]*Language = make(map[EnumLanguage]*Language)
)

func newLanguageEnum(key EnumLanguage) EnumLanguage {
	Languages = append(Languages, key)
	return EnumLanguage(key)
}

var (
	LanguageBrazilian = newLanguageEnum("pt_BR")
)

func loadLanguage(lang EnumLanguage) error {

	fileName := "internal/i18n/languages/" + string(lang) + ".yml"

	data, err := os.ReadFile(fileName)

	if err != nil {
		return err
	}

	var language Language

	err = yaml.Unmarshal(data, &language)

	if err != nil {
		return err
	}

	languageMap[lang] = &language

	return nil
}

func init() {

	for _, lang := range Languages {
		if err := loadLanguage(lang); err != nil {
			logger.Fatal("Failed to load languages: ", err)
		}
	}

}

func GetLanguage(lang EnumLanguage) *Language {
	language, found := languageMap[lang]

	if !found {
		language = languageMap[LanguageBrazilian]
	}

	return language
}

func Replace(str string, args ...interface{}) string {

	for i, value := range args {
		str = strings.Replace(str, fmt.Sprintf("${%d}", i), fmt.Sprintf("%v", value), 1)
	}

	return str

}
