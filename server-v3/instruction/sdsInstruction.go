package instruction

import (
	"fmt"
	"server-v3/global"
	"server-v3/model"
	"strconv"
)

func addSdsInsToMap() {
	InsMap["set"] = set()
	InsMap["get"] = get()
	InsMap["strlen"] = strlen()
	InsMap["del"] = del()
	InsMap["exists"] = exists()

	InsMap["mset"] = mset()
	InsMap["mget"] = mget()

	InsMap["incr"] = incr()
	InsMap["decr"] = decr()
	InsMap["incrby"] = incrby()
	InsMap["decrby"] = decrby()
}

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
			// check the length of the value
			if len(instruction.Value) == 1 {
				global.Map.SdsMap[instruction.Key] = model.Sds(instruction.Value[0])
				fmt.Println("OK")
			} else {
				fmt.Println("Failed: the length of the value should be 1")
			}
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
			printValueFromSdsMap(instruction.Key)
			return nil
		},
	}
}

func strlen() *model.Instruction {
	return &model.Instruction{
		Cmd: model.Command{
			Value:          "strlen",
			BelongTo:       CmdTypeEnum.Sds,
			WillChangeData: false,
		},
		Key:   "",
		Value: make([]string, 0),
		Execute: func(instruction *model.Instruction) error {
			fmt.Println(len(global.Map.SdsMap[instruction.Key]))
			return nil
		},
	}
}

func del() *model.Instruction {
	return &model.Instruction{
		Cmd: model.Command{
			Value:          "del",
			BelongTo:       CmdTypeEnum.Sds,
			WillChangeData: true,
		},
		Key:   "",
		Value: make([]string, 0),
		Execute: func(instruction *model.Instruction) error {
			delete(global.Map.SdsMap, instruction.Key)
			fmt.Println("OK")
			return nil
		},
	}
}

func exists() *model.Instruction {
	return &model.Instruction{
		Cmd: model.Command{
			Value:          "exists",
			BelongTo:       CmdTypeEnum.Sds,
			WillChangeData: false,
		},
		Key:   "",
		Value: make([]string, 0),
		Execute: func(instruction *model.Instruction) error {
			return nil
		},
	}
}

func mset() *model.Instruction {
	return &model.Instruction{
		Cmd: model.Command{
			Value:          "mset",
			BelongTo:       CmdTypeEnum.Sds,
			WillChangeData: true,
		},
		Key:   "",
		Value: make([]string, 0),
		Execute: func(instruction *model.Instruction) error {
			var kvArr = make([][2]string, 0, 2)
			// check the length of the value
			if len(instruction.Value)%2 == 0 {
				fmt.Println("Failed: the length of the value should be odd")
			} else {
				// store the key-value pair
				kvArr = append(kvArr, [2]string{instruction.Key, instruction.Value[0]})
				for i := 1; i < len(instruction.Value); i += 2 {
					kvArr = append(kvArr, [2]string{instruction.Value[i], instruction.Value[i+1]})
				}
				// store the key-value pair into the map
				for _, kv := range kvArr {
					global.Map.SdsMap[kv[0]] = model.Sds(kv[1])
				}
				fmt.Println("OK")
			}
			return nil
		},
	}
}

func mget() *model.Instruction {
	return &model.Instruction{
		Cmd: model.Command{
			Value:          "mget",
			BelongTo:       CmdTypeEnum.Sds,
			WillChangeData: false,
		},
		Key:   "",
		Value: make([]string, 0),
		Execute: func(instruction *model.Instruction) error {
			printValueFromSdsMap(instruction.Key)
			for _, v := range instruction.Value {
				printValueFromSdsMap(v)
			}
			return nil
		},
	}
}

func incr() *model.Instruction {
	return &model.Instruction{
		Cmd: model.Command{
			Value:          "incr",
			BelongTo:       CmdTypeEnum.Sds,
			WillChangeData: true,
		},
		Key:   "",
		Value: make([]string, 0),
		Execute: func(instruction *model.Instruction) error {
			commonIncr(instruction.Key, 1)
			return nil
		},
	}
}

func decr() *model.Instruction {
	return &model.Instruction{
		Cmd: model.Command{
			Value:          "decr",
			BelongTo:       CmdTypeEnum.Sds,
			WillChangeData: true,
		},
		Key:   "",
		Value: make([]string, 0),
		Execute: func(instruction *model.Instruction) error {
			commonIncr(instruction.Key, -1)
			return nil
		},
	}
}

func incrby() *model.Instruction {
	return &model.Instruction{
		Cmd: model.Command{
			Value:          "incrby",
			BelongTo:       CmdTypeEnum.Sds,
			WillChangeData: true,
		},
		Key:   "",
		Value: make([]string, 1),
		Execute: func(instruction *model.Instruction) error {
			// check the length of the value
			if len(instruction.Value) != 1 {
				fmt.Println("Failed: the length of the value should be 1")
			}
			// convert the value to int
			step, err := strconv.Atoi(instruction.Value[0])
			if err != nil {
				fmt.Println("Failed: the value is not an integer")
			} else {
				commonIncr(instruction.Key, step)
			}
			return nil
		},
	}
}

func decrby() *model.Instruction {
	return &model.Instruction{
		Cmd: model.Command{
			Value:          "decrby",
			BelongTo:       CmdTypeEnum.Sds,
			WillChangeData: true,
		},
		Key:   "",
		Value: make([]string, 1),
		Execute: func(instruction *model.Instruction) error {
			// check the length of the value
			if len(instruction.Value) != 1 {
				fmt.Println("Failed: the length of the value should be 1")
			}
			// convert the value to int
			step, err := strconv.Atoi(instruction.Value[0])
			if err != nil {
				fmt.Println("Failed: the value is not an integer")
			} else {
				commonIncr(instruction.Key, -step)
			}
			return nil
		},
	}
}

func printValueFromSdsMap(str string) {
	if v, ok := global.Map.SdsMap[str]; ok {
		fmt.Println(v)
	} else {
		fmt.Println("nil")
	}
}

func commonIncr(key string, step int) {
	// get value from the map
	v, ok := global.Map.SdsMap[key]
	if !ok { // if the key does not exist, then set the value to 1
		global.Map.SdsMap[key] = model.Sds(strconv.Itoa(step))
		fmt.Println(strconv.Itoa(step))
	} else { // if the key exists, then increment the value
		// convert the value to int
		i, err := strconv.Atoi(string(v))
		if err != nil { // if the value is not an integer, then print the error message
			fmt.Println("Failed: the value is not an integer")
		} else {
			i += step
			global.Map.SdsMap[key] = model.Sds(fmt.Sprintf("%d", i))
			fmt.Println(i)
		}
	}
}
