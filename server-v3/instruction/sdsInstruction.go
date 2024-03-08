package instruction

import (
	"fmt"
	"server-v3/global"
	"server-v3/model"
)

func set() *model.Instruction {
	return &model.Instruction{
		Cmd: model.Command{
			Value:          "set",
			BelongTo:       CmdTypeEnum.Sds,
			WillChangeData: true,
		},
		Key:   "",
		Value: make([]string, 1),
		Execute: func(instruction *model.Instruction) error {
			global.Map.SdsMap[instruction.Key] = model.Sds(instruction.Value[0])
			return nil
		},
	}
}

func get() *model.Instruction {
	return &model.Instruction{
		Cmd: model.Command{
			Value:          "get",
			BelongTo:       CmdTypeEnum.Sds,
			WillChangeData: false,
		},
		Key:   "",
		Value: make([]string, 0),
		Execute: func(instruction *model.Instruction) error {
			fmt.Println(global.Map.SdsMap[instruction.Key])
			return nil
		},
	}
}
