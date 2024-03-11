package tool

import (
	"server-v3/global"
	"strings"
)

type stringToolInterface interface {
	// TrimEndN trims the newline character from the end of the string.
	TrimEndN(text string) string
}

type stringTool struct {
}

func (s stringTool) TrimEndN(text string) string {
	switch global.Config.Os {
	case "windows":
		return strings.TrimSuffix(text, "\r\n")
	case "linux":
		return strings.TrimSuffix(text, "\n")
	case "mac":
		return strings.TrimSuffix(text, "\r")
	default:
		return text
	}
}

func (s stringTool) HandleUserInstruction(str string) string {
	// remove the newline character from the end of the string
	str = StringTool.TrimEndN(str)
	// remove the space character from the start and end of the string
	str = strings.TrimSpace(str)
	// replaces multiple Spaces in a string with a single space
	str = strings.Join(strings.Fields(str), " ")
	return str
}
