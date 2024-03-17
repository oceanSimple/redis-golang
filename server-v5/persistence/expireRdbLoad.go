package persistence

import (
	"bufio"
	"fmt"
	"go.uber.org/zap"
	"server-v5/global"
	"server-v5/log"
	"server-v5/tool"
	"strings"
	"time"
)

func ExpireRdbLoad() {
	fileConnector, err := tool.GetFileConnector("./persistence/expire_rdb.rdb")
	if err != nil {
		return
	}
	var loc, _ = time.LoadLocation("Local")
	scanner := bufio.NewScanner(fileConnector)
	for scanner.Scan() {
		split := strings.Split(scanner.Text(), ",")
		global.ExpireMap[split[0]], err = time.ParseInLocation("2006-01-02 15:04:05", split[1], loc)
		if err != nil {
			fmt.Println("expire rdb file read failed")
			log.SystemLog.Fatal("parse expire rdb time to time.Time failed",
				zap.Error(err))
		}
	}
	if err := scanner.Err(); err != nil {
		fmt.Println("aof file read failed")
		log.SystemLog.Fatal("aof file read failed",
			zap.Error(err))
	}
}
