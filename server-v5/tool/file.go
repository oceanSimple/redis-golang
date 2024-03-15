package tool

import (
	"go.uber.org/zap"
	"os"
	"server-v5/log"
)

func GetFileConnector(path string) (*os.File, error) {
	var fileConnector, err = os.OpenFile(path,
		os.O_APPEND|os.O_CREATE|os.O_RDWR, 0666)
	if err != nil {
		log.SystemLog.Error("Open file error, file path: "+path,
			zap.Error(err))
	}
	return fileConnector, err
}

func CloseFileConnector(fileConnector *os.File) {
	err := fileConnector.Close()
	if err != nil {
		log.SystemLog.Error("Close file error",
			zap.Error(err))
	}
}

func WriteToFile(fileConnector *os.File, str string) {
	_, err := fileConnector.WriteString(str)
	if err != nil {
		log.SystemLog.Error("Write to file error",
			zap.Error(err))
	}
}
