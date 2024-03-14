package persistence

import (
	"bufio"
	"fmt"
	"go.uber.org/zap"
	"server-v5/instruction"
	"server-v5/log"
	"server-v5/tool"
)

func LoadAofFile() {
	fileConnector, err := tool.GetFileConnector("./persistence/aof.aof")
	if err != nil {
		return
	}
	defer tool.CloseFileConnector(fileConnector)

	scanner := bufio.NewScanner(fileConnector)
	for scanner.Scan() {
		instruction.ExecuteInstruction(scanner.Text(), 0)
	}
	if err := scanner.Err(); err != nil {
		fmt.Println("aof file read failed")
		log.SystemLog.Fatal("aof file read failed",
			zap.Error(err))
	}
}
