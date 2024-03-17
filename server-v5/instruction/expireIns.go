package instruction

import (
	"server-v5/enum"
	"server-v5/global"
	"server-v5/model"
	"server-v5/persistence/rdb"
	"strconv"
	"time"
)

func addExpireInsToMap() {
	InsMap["expire"] = expire()
	InsMap["ttl"] = ttl()
}

// ttl
// return
// -2: the key is not exist in the whole maps, or the key is expired
// -1: the key is not in expire map
// >=0: the remaining time of the key
func ttl() *model.Instruction {
	return &model.Instruction{
		Cmd: model.Command{
			Value:          "ttl",
			BelongTo:       enum.Commands.Expire,
			WillChangeData: false,
		},
		Key:   "",
		Value: make([]string, 0),
		Execute: func(instruction *model.Instruction) (response []string) {
			response = append(response, strconv.Itoa(checkKeyIsExpired(instruction.Key)))
			return
		},
	}
}

// expire
// return
// "OK" if the expire time is set successfully
// "0" if the key is not existed
func expire() *model.Instruction {
	return &model.Instruction{
		Cmd: model.Command{
			Value:          "expire",
			BelongTo:       enum.Commands.Expire,
			WillChangeData: false,
		},
		Key:   "",
		Value: make([]string, 1),
		Execute: func(instruction *model.Instruction) (response []string) {
			// Check the value is only 1 and is int
			expireTime, str, ok := checkValueIsOnly1AndIsInt(instruction.Value)
			if !ok {
				response = append(response, str)
				return
			}
			// Query for mutex
			global.ReadMutexExpireMap.Lock()
			defer global.ReadMutexExpireMap.Unlock()
			// Check the key is existed
			if _, ok := global.Map[instruction.Key]; !ok {
				response = append(response, "0")
				return
			}
			// set the expire time
			global.ExpireMap[instruction.Key] = time.Now().Add(time.Duration(expireTime) * time.Second)
			// set the expire rdb
			rdb.ExpireMapRDB(instruction.Key, global.ExpireMap[instruction.Key])
			response = append(response, "1")
			return
		},
	}
}
