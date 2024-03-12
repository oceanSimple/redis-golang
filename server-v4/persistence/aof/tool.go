package aof

import (
	"go.uber.org/zap"
	"server-v4/log"
)

func writeToFile(instruction string) {
	_, err := fileConnector.WriteString(instruction)
	if err != nil {
		log.SystemLog.Fatal("aof file write failed",
			zap.Error(err))
	}
}
