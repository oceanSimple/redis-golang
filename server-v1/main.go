package main

import (
	"bufio"
	"fmt"
	"os"
	"server-v1/command"
)

func main() {
	terminalReader := bufio.NewReader(os.Stdin)
	var input string
	var cmd *command.Command
	var cmdTool = &command.Command{}
	for {
		fmt.Println("请输入命令:")
		input, _ = terminalReader.ReadString('\n')
		cmd = cmdTool.ParseCommand(input)
		_ = cmdTool.ExecuteCommand(cmd)
	}
}
