package instruction

import (
	"server-v4/global"
	"server-v4/model"
	"strconv"
	"time"
)

func setValueToSdsMap(key string, value string) {
	global.WholeMap.SdsMap[key] = model.Sds(value)
}

func getValueFromSdsMap(key string) string {
	var v, ok = global.WholeMap.SdsMap[key]
	if !ok {
		return "nil"
	} else {
		return string(v)
	}
}

// checkKeyIsExist
// check the key is existed in the whole map
func checkKeyIsExist(key string) bool {
	return checkKeyInMap(key, global.WholeMap.SdsMap) ||
		checkKeyInMap(key, global.WholeMap.ListMap) ||
		checkKeyInMap(key, global.WholeMap.HashMap) ||
		checkKeyInMap(key, global.WholeMap.SetMap)
}

// checkKeyInMap
// check the key is existed in the appointed map
// this function is a tool function for checkKeyIsExist
func checkKeyInMap[V any, T map[string]V](key string, m T) bool {
	_, ok := m[key]
	return ok
}

// checkValueLen1AndIsInt
// check the value is a slice with length 1 and the value is an integer
func checkValueLen1AndIsInt(value []string) (int, string, bool) {
	if len(value) != 1 {
		return 0, "the length of value must be 1", false
	}
	i, err := strconv.Atoi(value[0])
	if err != nil {
		return 0, "The value is not an integer", false
	}
	return i, "", true
}

// checkKeyIsExpired
// check the key is expired or not
// return the remaining time of the key
// -2: the key is not exist in the whole maps, or the key is expired
// -1: the key is not in expire map
// >=0: the remaining time of the key
func checkKeyIsExpired(key string) int {
	// Condition 1: key is not exist
	if existed := checkKeyIsExist(key); !existed {
		return -2
	}
	// Condition 2: key is exist, but not in expire map
	if existed := checkKeyInMap(key, global.WholeMap.ExpireMap); !existed {
		return -1
	}
	// Condition 3: key is exist, and in expire map
	ttl := global.WholeMap.ExpireMap[key].Sub(time.Now())
	if ttl >= 0 {
		return int(ttl.Seconds())
	} else {
		return -2
	}
}

// mapNeedDelKey
// check the key is needed to be deleted
// return true if the key is deleted
// return false if the key is not deleted
func mapNeedDelKey[V any, T map[string]V](key string, m T) bool {
	code := checkKeyIsExpired(key)
	if code == -2 {
		delete(m, key)                         // delete the key in target map
		delete(global.WholeMap.ExpireMap, key) // delete the key in expire map
		return true
	}
	return false
}
