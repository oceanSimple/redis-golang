package main

import (
	"bufio"
	"fmt"
	"os"
	"server-v5/instruction"
	"server-v5/persistence"
)

func main() {
	aofLoad()
	instructionWhile()
}

func instructionWhile() {
	var (
		reader = bufio.NewReader(os.Stdin)
		cmdStr string
		err    error
	)
	for {
		fmt.Print("> ")
		cmdStr, err = reader.ReadString('\n')
		if err != nil {
			break
		}
		// cmdStr = tool.StringTool.HandleUserInstruction(cmdStr)
		instruction.ExecuteInstruction(cmdStr, 1)
	}
}

func aofLoad() {
	persistence.LoadAofFile()
}
