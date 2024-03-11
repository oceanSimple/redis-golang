package instruction

import (
	"go.uber.org/zap"
	"server-v3/log"
	"server-v3/persistence/aof"
	"strings"
)

// ExecuteInstruction flag is the flag of the instruction
// 0: system --- aof init memory
// 1: user --- user input
func ExecuteInstruction(str string, flag int) {
	splits := strings.Split(str, " ")
	// TODO Now, we only support the command which has at least 2 parts
	if len(splits) < 2 {
		return
	}
	// get the instruction from the map
	ins := InsMap[splits[0]]
	if ins == nil {
		return
	}
	ins.Key = splits[1]
	ins.Value = splits[2:]

	err := ins.Execute(ins)
	if err != nil {
		log.SystemLog.Error("failed to execute instruction",
			zap.Error(err))
		return
	}

	// If the command will change data, then enter the aof mode
	if ins.Cmd.WillChangeData && flag != 0 {
		aof.InsEntrance(str + "\n")
	}
}
