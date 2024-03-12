package main

import (
	"bufio"
	"fmt"
	"os"
	"server-v4/instruction"
	"server-v4/persistence"
	"server-v4/persistence/aof"
)

func main() {
	LoadData()
	startAofRoutine()
	instructionWhile()
}

func LoadData() {
	persistence.LoadAofFile()
}

func startAofRoutine() {
	go func() {
		aof.GoRoutineMethod()
	}()
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
