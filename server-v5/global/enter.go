package global

import (
	"server-v5/model"
	"sync"
	"time"
)

// Export the global variables

var (
	// Map store all the map
	Map map[string]*model.RedisObject
	// RehashMap Reserved for later rehash
	RehashMap map[string]*model.RedisObject
	// ExpireMap store the expire key and the expire time
	ExpireMap map[string]time.Time
)

var (
	WriteMutexMup = sync.Mutex{}
	ReadMutexMap  = sync.RWMutex{}
)

var (
	AofRewriteFlag bool
)
