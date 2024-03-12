package instruction

import (
	"fmt"
	"server-v3/global"
	"server-v3/model"
	"strconv"
	"time"
)

func addExpireInsToMap() {
	InsMap["expire"] = expire()
	InsMap["ttl"] = ttl()
}

func ttl() *model.Instruction {
	return &model.Instruction{
		Cmd: model.Command{
			Value:          "ttl",
			BelongTo:       CmdTypeEnum.Expire,
			WillChangeData: false,
		},
		Key:   "",
		Value: make([]string, 0),
		Execute: func(instruction *model.Instruction) error {
			// look up the expire map
			if v, ok := global.Map.ExpireMap[instruction.Key]; ok {
				ttl := v.Sub(time.Now())
				if ttl > 0 {
					fmt.Println(ttl)
				} else {
					fmt.Println(-2)
				}
			} else {
				fmt.Println(-1)
			}
			return nil
		},
	}
}

func expire() *model.Instruction {
	return &model.Instruction{
		Cmd: model.Command{
			Value:          "expire",
			BelongTo:       CmdTypeEnum.Expire,
			WillChangeData: true,
		},
		Key:   "",
		Value: make([]string, 0),
		Execute: func(instruction *model.Instruction) error {
			// check the value
			if len(instruction.Value) != 1 {
				fmt.Println("The value is not empty")
				return nil
			}
			expireTime, err := strconv.Atoi(instruction.Value[0])
			if err != nil {
				fmt.Println("The value is not an integer")
				return nil
			}

			// check the key is existed

			// look up the expire map
			if _, ok := global.Map.ExpireMap[instruction.Key]; !ok {
				global.Map.ExpireMap[instruction.Key] = time.Now().Add(time.Duration(expireTime) * time.Second)
				fmt.Println("OK")
			} else {
				fmt.Println("The key does not exist")
			}
			return nil
		},
	}
}
