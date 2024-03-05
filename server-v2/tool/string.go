package tool

import (
	"server-v2/global"
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
