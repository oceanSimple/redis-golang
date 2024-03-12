package global

import (
	"container/list"
	"server-v4/model"
	"time"
)

type MapSet struct {
	SdsMap    map[string]model.Sds
	ListMap   map[string][]list.List
	HashMap   map[string]map[string]model.Sds
	SetMap    map[string]model.Set
	ExpireMap map[string]time.Time
}
