package instruction

import "server-v5/model"

func init() {
	instructionMapInit()
}

func instructionMapInit() {
	InsMap = make(map[string]*model.Instruction)
	addExpireInsToMap()
	addSdsInsToMap()
}
