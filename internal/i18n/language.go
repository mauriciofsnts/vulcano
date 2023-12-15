package i18n

type LanguageEntry string

func (l LanguageEntry) Str(args ...interface{}) string {
	return Replace(string(l), args...)
}

type LanguageMetadata struct {
	Name      LanguageEntry
	ShortName LanguageEntry
	Author    LanguageEntry
}

type Errors struct {
	Title             LanguageEntry
	Generic           LanguageEntry
	ToSave            LanguageEntry
	NotATextChannel   LanguageEntry
	AlreadyRegistered LanguageEntry
	MissingParamter   LanguageEntry
	InvalidParameter  LanguageEntry
}

type Command struct {
	Title       LanguageEntry
	Response    LanguageEntry
	Description LanguageEntry
}

type Commands struct {
	Ping                Command
	Tools               Command
	Tabnews             Command
	Shorten             Command
	Holiday             Command
	Uptime              Command
	Help                Command
	Generate            Command
	OpenSourcesProjects Command
}

type Language struct {
	Lang     LanguageMetadata
	Errors   Errors
	Commands Commands
	Utils    Utils
}

type Uptime struct {
	Title       LanguageEntry
	Description LanguageEntry
	Response    LanguageEntry
}

type Utils struct {
	LessThanAMinute LanguageEntry
	Minutes         LanguageEntry
	Minute          LanguageEntry
	Hours           LanguageEntry
	Hour            LanguageEntry
	Days            LanguageEntry
	Day             LanguageEntry
	Months          LanguageEntry
	Month           LanguageEntry
	Years           LanguageEntry
	Year            LanguageEntry
}

type Help struct {
	Description LanguageEntry
	Title       LanguageEntry
	Response    LanguageEntry
}
