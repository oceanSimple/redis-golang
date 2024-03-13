package global

import (
	"container/list"
	"runtime"
	"server-v4/model"
	"time"
)

var (
	// WholeMap store all the map
	WholeMap MapSet
	// Config store all the system config
	Config RunTimeConfig
)

func init() {
	mapInit()
	configInit()
}

func mapInit() {
	WholeMap = MapSet{
		SdsMap:    make(map[string]model.Sds),
		ListMap:   make(map[string][]list.List),
		HashMap:   make(map[string]map[string]model.Sds),
		SetMap:    make(map[string]model.Set),
		ExpireMap: make(map[string]time.Time),
	}
}

func configInit() {
	Config = RunTimeConfig{
		Os: runtime.GOOS,
	}
}
