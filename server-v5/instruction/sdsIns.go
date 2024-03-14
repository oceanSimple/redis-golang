package instruction

import (
	"server-v5/enum"
	"server-v5/global"
	"server-v5/model"
)

func addSdsInsToMap() {
	InsMap["set"] = set()
	InsMap["get"] = get()
	InsMap["del"] = del()
}

func set() *model.Instruction {
	return &model.Instruction{
		Cmd: model.Command{
			Value:          "set",
			BelongTo:       enum.Commands.String,
			WillChangeData: true,
		},
		Key:   "",
		Value: make([]string, 1),
		Execute: func(instruction *model.Instruction) (response []string) {
			// check the length of the value
			if len(instruction.Value) == 1 {
				setNewKeyToMap(instruction.Key, enum.Types.String, instruction.Value[0])
				response = append(response, "OK")
			} else {
				response = append(response, "Failed: the length of the value should be 1")
			}
			return
		},
	}
}

func get() *model.Instruction {
	return &model.Instruction{
		Cmd: model.Command{
			Value:          "get",
			BelongTo:       enum.Commands.String,
			WillChangeData: false,
		},
		Key:   "",
		Value: make([]string, 0),
		Execute: func(instruction *model.Instruction) (response []string) {
			response = append(response, getFromMap(instruction.Key))
			return
		},
	}
}

// del
// TODO now the del function is not completed, it should delete multiple keys
// TODO ,and return the number of keys deleted
func del() *model.Instruction {
	return &model.Instruction{
		Cmd: model.Command{
			Value:          "del",
			BelongTo:       enum.Commands.String,
			WillChangeData: true,
		},
		Key:   "",
		Value: make([]string, 0),
		Execute: func(instruction *model.Instruction) (response []string) {
			delete(global.Map, instruction.Key)
			response = append(response, "OK")
			return
		},
	}
}
