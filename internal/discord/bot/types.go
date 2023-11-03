package bot

var (
	TypeString = &BaseType{Name: "string"}
	TypeInt    = &BaseType{Name: "int"}
	TypeBool   = &BaseType{Name: "bool"}
)

const (
	CommandTriggerSlash   TriggerType = "SLASH"
	CommandTriggerMessage TriggerType = "MESSAGE"
)
