package instruction

import (
	"server-v3/persistence/aof"
	"strings"
)

// ExecuteInstruction flag is the flag of the instruction
// 0: system --- aof init memory
// 1: user --- user input
func ExecuteInstruction(str string, flag int) {
	splits := strings.Split(str, " ")
	ins := InsMap[splits[0]]
	if ins == nil {
		return
	}
	ins.Key = splits[1]
	ins.Value = splits[2:]

	err := ins.Execute(ins)
	if err != nil {
		return
	}

	// If the command will change data, then enter the aof mode
	if ins.Cmd.WillChangeData && flag != 0 {
		aof.InsEntrance(str + "\n")
	}
}
