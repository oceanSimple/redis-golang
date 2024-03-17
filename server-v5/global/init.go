package global

import (
	"server-v5/model"
	"time"
)

func init() {
	Map = make(map[string]*model.RedisObject)
	RehashMap = make(map[string]*model.RedisObject)
	ExpireMap = make(map[string]time.Time)
}
