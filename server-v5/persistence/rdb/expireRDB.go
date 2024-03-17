package rdb

import (
	"fmt"
	"go.uber.org/zap"
	"server-v5/log"
	"server-v5/tool"
	"time"
)

func ExpireMapRDB(key string, expireTime time.Time) {
	fileConnector, err := tool.GetFileConnector("./persistence/expire_rdb.rdb")
	if err != nil {
		panic(err)
	}
	defer tool.CloseFileConnector(fileConnector)

	_, err = fileConnector.WriteString(fmt.Sprintf("%s,%s\n", key, expireTime.Format("2006-01-02 15:04:05")))
	if err != nil {
		log.SystemLog.Error("write expire rdb error: ",
			zap.Error(err))
	}
}
