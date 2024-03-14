package instruction

import (
	"fmt"
	"server-v5/global"
	"server-v5/model"
	"server-v5/persistence/aof"
	"time"
)

// setNewKeyToMap
// set a new key and value to the map
func setNewKeyToMap(key string, typeName string, value any) string {
	// Check if the key already exists,
	// and the type of the key is not the same as the new type
	if v, ok := global.Map[key]; ok && v.Type != typeName {
		return "Operation against a key holding the wrong kind of value"
	}
	global.Map[key] = &model.RedisObject{
		Type: typeName,
		Ptr:  value,
	}
	return "OK"
}

// getFromMap
// get the value from the map
func getFromMap(key string) string {
	// Check the key is existed or not
	v, ok := global.Map[key]
	if !ok {
		return "nil"
	}
	// Check if the key is expired
	var ttl = checkKeyIsExpired(key)
	if ttl == -2 {
		delete(global.Map, key)
		delete(global.ExpireMap, key)
		// write it to the log
		aof.WriteToAof("del " + key)
		return "nil"
	}
	return fmt.Sprintf("%v", v.Ptr)
}

// checkKeyIsExpired
// check the key is expired or not
// return the remaining time of the key
// -2: the key is not exist in the whole maps, or the key is expired
// -1: the key is not in expire map
// >=0: the remaining time of the key
func checkKeyIsExpired(key string) int {
	// Condition 1: key is not exist
	if _, ok := global.Map[key]; !ok {
		return -2
	}
	// Condition 2: key is not in expire map
	if _, ok := global.ExpireMap[key]; !ok {
		return -1
	}
	// Condition 3: key is expired
	if global.ExpireMap[key].Before(time.Now()) {
		return -2
	}
	// Condition 4: key is not expired
	return int(global.ExpireMap[key].Sub(time.Now()).Seconds())
}
