package instruction

import (
	"server-v4/global"
	"server-v4/model"
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
			BelongTo:       cmdTypeEnum.Expire,
			WillChangeData: false,
		},
		Key:   "",
		Value: make([]string, 0),
		Execute: func(instruction *model.Instruction) ([]string, error) {
			var response = make([]string, 0, 1)
			response = append(response, strconv.Itoa(checkKeyIsExpired(instruction.Key)))
			return response, nil
		},
	}
}

func expire() *model.Instruction {
	return &model.Instruction{
		Cmd: model.Command{
			Value:          "expire",
			BelongTo:       cmdTypeEnum.Expire,
			WillChangeData: false,
		},
		Key:   "",
		Value: make([]string, 1),
		Execute: func(instruction *model.Instruction) ([]string, error) {
			var response = make([]string, 0, 1)
			// check the value
			expireTime, str, ok := checkValueLen1AndIsInt(instruction.Value)
			if !ok {
				response = append(response, str)
				return response, nil
			}
			// check the key is existed
			if existed := checkKeyIsExist(instruction.Key); !existed {
				response = append(response, "The key is not exist")
				return response, nil
			}
			// set the expire time
			global.WholeMap.ExpireMap[instruction.Key] = time.Now().Add(time.Duration(expireTime) * time.Second)
			response = append(response, "OK")
			return response, nil
		},
	}
}
