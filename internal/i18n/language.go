package i18n

type LanguageEntry string

func (l LanguageEntry) Str(args ...interface{}) string {
	return Replace(string(l), args...)
}

type Language struct {
	Commands Commands
	Intro    string
}

type Command struct {
	Name        LanguageEntry
	Description LanguageEntry
	Reply       LanguageEntry
	Error       LanguageEntry
}

type Commands struct {
	Devto    Command
	Tabnews  Command
	Newsapi  Command
	Balance  Command
	Uptime   Command
	Help     Command
	Services Command
	Shorten  Command
	Ping     Command
	Generate Command
}
