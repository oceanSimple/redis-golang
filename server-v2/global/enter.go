package global

import (
	"runtime"
	"server-v2/model"
)

var (
	Map     MapSet
	Config  RunTimeConfig
	CmdType *CmdTypeEnum
	CmdMap  map[string]string
)

func init() {
	mapInit()
	configInit()
	cmdTypeInit()
	cmdMapInit()
}

func mapInit() {
	Map = MapSet{
		SdsMap: make(map[string]model.Sds),
	}
	Map.SdsMap["test"] = model.Sds("test")
}

func configInit() {
	Config = RunTimeConfig{
		Os: runtime.GOOS,
	}
}

func cmdMapInit() {
	CmdMap = make(map[string]string)
	addExpireCommand(CmdMap)
	addSdsCommand(CmdMap)
	addListCommand(CmdMap)
	addHashCommand(CmdMap)
	addSetCommand(CmdMap)
}

func cmdTypeInit() {
	CmdType = &CmdTypeEnum{}
	CmdType.init()
}
