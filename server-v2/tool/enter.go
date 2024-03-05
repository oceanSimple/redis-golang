package tool

var (
	CmdTool commandTool
	StrTool stringTool
)

func init() {
	CmdTool = commandTool{}
	StrTool = stringTool{}
}
