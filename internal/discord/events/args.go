package events

import "strconv"

func (cmnd CommandMessage) GetArgString(idx int) string {
	if cmnd.Interaction != nil {
		args := cmnd.Interaction.Args
		return args[idx].StringValue()
	} else {
		args := cmnd.Message.Args
		return args[idx]
	}
}

func (cmnd CommandMessage) GetArgInt(idx int) (int, error) {
	if cmnd.Interaction != nil {
		args := cmnd.Interaction.Args
		return int(args[idx].IntValue()), nil
	} else {
		args := cmnd.Message.Args
		value, err := strconv.ParseInt(args[idx], 10, 64)
		return int(value), err
	}
}

func (cmnd CommandMessage) HasArg(idx int) bool {
	if cmnd.Interaction != nil {
		args := cmnd.Interaction.Args
		return len(args) > idx
	} else {
		args := cmnd.Message.Args
		return len(args) > idx
	}
}
