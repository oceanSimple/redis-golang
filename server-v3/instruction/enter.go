package instruction

import "server-v3/model"

var (
	CmdTypeEnum *CommandTypeEnum
	InsMap      map[string]*model.Instruction
)

func init() {
	cmdTypeEnumInit()
	instructionMapInit()
}

func cmdTypeEnumInit() {
	CmdTypeEnum = &CommandTypeEnum{}
	CmdTypeEnum.init()
}

func instructionMapInit() {
	InsMap = make(map[string]*model.Instruction)
	addSdsInsToMap()
}
