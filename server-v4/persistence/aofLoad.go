package persistence

import (
	"bufio"
	"go.uber.org/zap"
	"os"
	"server-v4/instruction"
	"server-v4/log"
)

func LoadAofFile() {
	fileConnector, err := os.OpenFile("./persistence/aof.aof",
		os.O_CREATE|os.O_APPEND|os.O_RDWR,
		0666)
	defer func(fileConnector *os.File) {
		err := fileConnector.Close()
		if err != nil {
			log.SystemLog.Fatal("Aof loading: aof file close failed",
				zap.Error(err))
		}
	}(fileConnector)
	if err != nil {
		log.SystemLog.Fatal("Aof loading: aof file open failed",
			zap.Error(err))
	}

	scanner := bufio.NewScanner(fileConnector)
	for scanner.Scan() {
		instruction.ExecuteInstruction(scanner.Text(), 0)
	}
	if err := scanner.Err(); err != nil {
		log.SystemLog.Fatal("aof file read failed",
			zap.Error(err))
	}
}
