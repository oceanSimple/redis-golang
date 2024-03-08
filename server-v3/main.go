package main

import (
	"bufio"
	"fmt"
	"os"
	"server-v3/instruction"
	"server-v3/tool"
)

func main() {
	var (
		reader = bufio.NewReader(os.Stdin)
		cmdStr string
		err    error
		// cmd    *model.Command
	)
	for {
		fmt.Print("> ")
		cmdStr, err = reader.ReadString('\n')
		if err != nil {
			break
		}
		cmdStr = tool.StringTool.TrimEndN(cmdStr)
		instruction.ExecuteInstruction(cmdStr)
	}
}
