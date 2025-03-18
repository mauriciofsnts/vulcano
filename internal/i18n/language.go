package i18n

type LanguageEntry string

func (l LanguageEntry) Str(args ...interface{}) string {
	return Replace(string(l), args...)
}

type Language struct {
	Commands Commands
	Intro    string
	Global   Global
}

type Command struct {
	Name        LanguageEntry
	Description LanguageEntry
	Reply       LanguageEntry
	Error       LanguageEntry
}

type Event struct {
	Name    LanguageEntry
	Error   LanguageEntry
	Victory LanguageEntry
	Defeat  LanguageEntry
}

type GenerateCommand struct {
	Command
	ParamError LanguageEntry `yaml:"paramError"`
	WithMask   LanguageEntry `yaml:"withMask"`
	Options    LanguageEntry `yaml:"options"`
}

type Commands struct {
	Devto       Command
	Tabnews     Command
	Newsapi     Command
	Balance     Command
	Uptime      Command
	Help        Command
	Services    Command
	Shorten     Command
	Ping        Command
	Brasileirao Command
	Generate    GenerateCommand
}

type Global struct {
	LessThatAMinute LanguageEntry `yaml:"less_than_a_minute"`
	Uptime          LanguageEntry
	Latency         LanguageEntry
	Database        LanguageEntry
}
