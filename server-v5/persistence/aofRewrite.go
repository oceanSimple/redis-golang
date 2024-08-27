package persistence

//func ReWriteAof() {
//	// Set the aof rewrite flag to true
//	valueMap.AofRewriteFlag = true
//	// Copy the map to a temp map
//	tempMap := make(map[string]*model.RedisObject)
//	valueMap.ReadMutexMap.RLock()
//	for k, v := range valueMap.Map {
//		tempMap[k] = v
//	}
//	valueMap.ReadMutexMap.RUnlock()
//
//	// Write the temp map to the aof file
//	fileConnector, err := tool.GetFileConnector("./persistence/temp_aof.aof")
//	defer tool.CloseFileConnector(fileConnector)
//	if err != nil {
//		log.SystemLog.Error("failed to open the aof rewrite file",
//			zap.Error(err))
//	}
//	for k, v := range tempMap {
//		s := fmt.Sprintf("SET %s %s\n", k, v.Ptr)
//		tool.WriteToFile(fileConnector, s)
//	}
//}
