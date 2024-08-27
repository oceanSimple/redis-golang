package valueMap

import (
	"server-v5-refactor-server/static/structure"
	"sync"
)

var (
	Store map[string]*structure.RedisObject
)

var (
	// WMutex  = sync.Mutex{}
	RWMutex = sync.RWMutex{}
)

func init() {
	Store = make(map[string]*structure.RedisObject)
}
