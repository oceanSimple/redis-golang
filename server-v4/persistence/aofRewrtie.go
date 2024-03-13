package persistence

import (
	"fmt"
	"go.uber.org/zap"
	"io"
	"os"
	"server-v4/global"
	"server-v4/log"
	"server-v4/persistence/aof"
)

type reWriteFuncInterface interface {
	Rewrite()
	openFile() *os.File
	coverOldFile()
}

type ReWriteFunc struct {
}

func (r ReWriteFunc) Rewrite() {
	fileConnector := r.openFile("./persistence/aof_temp.aof")
	defer func(fileConnector *os.File) {
		err := fileConnector.Close()
		if err != nil {
			log.SystemLog.Fatal("Aof rewrite: aof file close failed",
				zap.Error(err))
		}
	}(fileConnector)

	r.writeSdsMapToFile(fileConnector)

	r.coverOldFile(fileConnector)
}

func (r ReWriteFunc) openFile(path string) *os.File {
	// Open a temporary file to write the new aof file
	fileConnector, err := os.OpenFile(path,
		os.O_CREATE|os.O_RDWR|os.O_TRUNC,
		0666)
	if err != nil {
		log.SystemLog.Fatal("Aof rewrite: aof file open failed",
			zap.Error(err))
	}
	return fileConnector
}

func (r ReWriteFunc) writeSdsMapToFile(fileConnector *os.File) {
	for k, v := range global.WholeMap.SdsMap {
		ins := fmt.Sprintf("set %s %s\n", k, v)
		_, err := fileConnector.WriteString(ins)
		if err != nil {
			log.SystemLog.Fatal("Aof rewrite: write sds map to file failed",
				zap.Error(err))
		}
	}
}

func (r ReWriteFunc) coverOldFile(fileConnector *os.File) {
	// Cover the old aof file
	if _, err := io.Copy(aof.GetFileConnector(), fileConnector); err != nil {
		log.SystemLog.Fatal("Aof rewrite: cover old file failed",
			zap.Error(err))
	}

	// Sync the file to disk
	if err := fileConnector.Sync(); err != nil {
		log.SystemLog.Fatal("Aof rewrite: sync file failed",
			zap.Error(err))
	}
}
