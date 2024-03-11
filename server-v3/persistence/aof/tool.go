package aof

import (
	"go.uber.org/zap"
	"server-v3/log"
)

func writeToFile(instruction string) {
	_, err := FileConnector.WriteString(instruction)
	if err != nil {
		log.SystemLog.Fatal("aof file write failed",
			zap.Error(err))
	}
}
