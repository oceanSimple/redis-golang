package global

import (
	"server-v2/model"
)

var (
	Map MapSet
)

func init() {
	mapInit()
}

func mapInit() {
	Map = MapSet{
		SdsMap: make(map[string]model.Sds),
	}
	Map.SdsMap["test"] = model.Sds("test")
}
