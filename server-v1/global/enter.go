package global

import (
	"server-v1/model"
	"server-v1/service/serviceImpl"
)

var (
	GlobalBuilder *Builder
	GlobalMap     *MapSet
)

func init() {
	initGlobalBuilder()
	initGlobalMap()
}

func initGlobalBuilder() {
	GlobalBuilder = &Builder{
		SdsBuilder: &serviceImpl.SdsServiceImpl{},
	}
}

func initGlobalMap() {
	GlobalMap = &MapSet{
		SdsMap: make(map[string]*model.Sds),
	}
	GlobalMap.SdsMap["test"] = GlobalBuilder.SdsBuilder.NewSds("test")
}
