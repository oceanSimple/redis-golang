package tool

import "server-v2/command"

var (
	CmdTool command.Tool
)

func init() {
	CmdTool = command.Tool{}
}
