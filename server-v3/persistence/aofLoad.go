package persistence

import (
	"bufio"
	"go.uber.org/zap"
	"server-v3/instruction"
	"server-v3/log"
	"server-v3/persistence/aof"
)

func LoadAofFile() {
	scanner := bufio.NewScanner(aof.FileConnector)
	for scanner.Scan() {
		instruction.ExecuteInstruction(scanner.Text(), 0)
	}
	if err := scanner.Err(); err != nil {
		log.SystemLog.Fatal("aof file read failed",
			zap.Error(err))
	}
}
