package instruction

import (
	"server-v4/global"
	"server-v4/model"
	"strconv"
)

func addSdsInsToMap() {
	InsMap["set"] = set()
	InsMap["get"] = get()
	InsMap["strlen"] = strlen()
	InsMap["del"] = del()
}

func set() *model.Instruction {
	return &model.Instruction{
		Cmd: model.Command{
			Value:          "set",
			BelongTo:       cmdTypeEnum.Sds,
			WillChangeData: true,
		},
		Key:   "",
		Value: make([]string, 1),
		Execute: func(instruction *model.Instruction) ([]string, error) {
			var response = make([]string, 0, 1)
			// check the length of the value
			if len(instruction.Value) == 1 {
				setValueToSdsMap(instruction.Key, instruction.Value[0])
				response = append(response, "OK")
			} else {
				response = append(response, "Failed: the length of the value should be 1")
			}
			return response, nil
		},
	}
}

func get() *model.Instruction {
	return &model.Instruction{
		Cmd: model.Command{
			Value:          "get",
			BelongTo:       cmdTypeEnum.Sds,
			WillChangeData: false,
		},
		Key:   "",
		Value: make([]string, 0),
		Execute: func(instruction *model.Instruction) ([]string, error) {
			var response = make([]string, 0, 1)
			// check the key is expired, if expired, return nil
			if ok := mapNeedDelKey(instruction.Key, global.WholeMap.SdsMap); ok {
				response = append(response, "nil")
				return response, nil
			}
			response = append(response, getValueFromSdsMap(instruction.Key))
			return response, nil
		},
	}
}

func strlen() *model.Instruction {
	return &model.Instruction{
		Cmd: model.Command{
			Value:          "strlen",
			BelongTo:       cmdTypeEnum.Sds,
			WillChangeData: false,
		},
		Key:   "",
		Value: make([]string, 0),
		Execute: func(instruction *model.Instruction) ([]string, error) {
			var response = make([]string, 0, 1)
			response = append(response,
				strconv.Itoa(len(global.WholeMap.SdsMap[instruction.Key])))
			return response, nil
		},
	}
}

func del() *model.Instruction {
	return &model.Instruction{
		Cmd: model.Command{
			Value:          "del",
			BelongTo:       cmdTypeEnum.Sds,
			WillChangeData: true,
		},
		Key:   "",
		Value: make([]string, 0),
		Execute: func(instruction *model.Instruction) ([]string, error) {
			var response = make([]string, 0, 1)
			delete(global.WholeMap.SdsMap, instruction.Key)
			response = append(response, "OK")
			return response, nil
		},
	}
}
