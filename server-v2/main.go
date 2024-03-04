package main

import (
	"bufio"
	"fmt"
	"os"
	"server-v2/model"
	"server-v2/tool"
)

func main() {
	terminalReader := bufio.NewReader(os.Stdin)
	var input string
	var cmd *model.Command
	var cmdTool = tool.CmdTool
	for {
		fmt.Println("请输入命令:")
		input, _ = terminalReader.ReadString('\n')
		cmd = cmdTool.ParseCommand(input)
		_ = cmdTool.ExecuteCommand(cmd)
	}
}
