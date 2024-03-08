package global

import (
	"runtime"
	"server-v3/model"
)

var (
	Map    MapSet
	Config RunTimeConfig
)

func init() {
	mapInit()
	configInit()
}

func mapInit() {
	Map = MapSet{
		SdsMap: make(map[string]model.Sds),
	}
	Map.SdsMap["test"] = "test"
}

func configInit() {
	Config = RunTimeConfig{
		Os: runtime.GOOS,
	}
}
