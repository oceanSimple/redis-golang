package main

import (
	"bufio"
	"fmt"
	"os"
	_ "server-v5/config"
	"server-v5/instruction"
	"server-v5/persistence"
	"server-v5/persistence/aof"
)

func main() {
	aofLoad()
	expireRdbLoad()
	go aof.GoRoutineMethod()
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

func expireRdbLoad() {
	persistence.ExpireRdbLoad()
}
