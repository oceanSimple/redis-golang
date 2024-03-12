package instruction

import "server-v4/model"

func init() {
	cmdTypeEnumInit()
	instructionMapInit()
}

func cmdTypeEnumInit() {
	cmdTypeEnum = &commandTypeEnum{}
	cmdTypeEnum.init()
}

func instructionMapInit() {
	InsMap = make(map[string]*model.Instruction)
	addSdsInsToMap()
	addExpireInsToMap()
}
