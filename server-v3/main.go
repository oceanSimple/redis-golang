package main

import (
	"bufio"
	"fmt"
	"os"
	"server-v3/instruction"
	"server-v3/persistence"
	"server-v3/persistence/aof"
	"server-v3/tool"
	_ "server-v3/viper"
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
		cmdStr = tool.StringTool.HandleUserInstruction(cmdStr)
		instruction.ExecuteInstruction(cmdStr, 1)
	}
}
