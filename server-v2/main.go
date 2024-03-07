package main

import (
	"bufio"
	"os"
	"server-v2/command"
	"server-v2/tool"
)

func main() {
	var (
		reader = bufio.NewReader(os.Stdin)
		cmdStr string
		err    error
		// cmd    *model.Command
	)
	for {
		cmdStr, err = reader.ReadString('\n')
		if err != nil {
			break
		}
		cmdStr = tool.StrTool.TrimEndN(cmdStr)
		_ = command.CmdParseTool.ExecuteCmd(cmdStr)
	}
}
