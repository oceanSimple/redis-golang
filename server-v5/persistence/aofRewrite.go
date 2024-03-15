package persistence

// TODO Failed to finish this function
import (
	"fmt"
	"go.uber.org/zap"
	"server-v5/global"
	"server-v5/log"
	"server-v5/model"
	"server-v5/tool"
)

func ReWriteAof() {
	// Set the aof rewrite flag to true
	global.AofRewriteFlag = true
	// Copy the map to a temp map
	tempMap := make(map[string]*model.RedisObject)
	global.ReadMutexMap.RLock()
	for k, v := range global.Map {
		tempMap[k] = v
	}
	global.ReadMutexMap.RUnlock()

	// Write the temp map to the aof file
	fileConnector, err := tool.GetFileConnector("./persistence/temp_aof.aof")
	defer tool.CloseFileConnector(fileConnector)
	if err != nil {
		log.SystemLog.Error("failed to open the aof rewrite file",
			zap.Error(err))
	}
	for k, v := range tempMap {
		s := fmt.Sprintf("SET %s %s\n", k, v.Ptr)
		tool.WriteToFile(fileConnector, s)
	}
}
