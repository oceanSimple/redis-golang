package valueMap

import "server-v5-refactor-server/static/structure"

func GetValueByKey(key string) *structure.RedisObject {
	RWMutex.RLock()
	obj := Store[key]
	RWMutex.RUnlock()
	return obj
}

func SetValueByKey(key string, obj *structure.RedisObject) {
	RWMutex.Lock()
	Store[key] = obj
	RWMutex.Unlock()
}

func DelValueByKey(key string) {
	RWMutex.Lock()
	delete(Store, key)
	RWMutex.Unlock()
}
