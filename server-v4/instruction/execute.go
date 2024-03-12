package instruction

import (
	"fmt"
	"server-v4/persistence/aof"
	"strings"
)

// ExecuteInstruction flag is the flag of the instruction
// 0: system --- aof init memory
// 1: user --- user input
func ExecuteInstruction(str string, flag int) {
	// Split the string and check the length
	// Now we only support the command which has at least 2 parts
	splits := strings.Fields(str)
	if len(splits) < 2 {
		fmt.Println("ERR wrong number of arguments for 'execute' command")
		return
	}

	// Get the instruction from the map
	ins := InsMap[splits[0]]
	if ins == nil {
		fmt.Println("ERR unknown command '" + splits[0] + "'")
		return
	}

	// Fill key and value
	ins.Key = splits[1]
	ins.Value = splits[2:]
	// execute the instruction
	strArr, err := ins.Execute(ins)

	// Check the error
	if err != nil {
		fmt.Println("ERR " + err.Error())
		return
	}

	// print the result
	if flag != 0 {
		for _, s := range strArr {
			fmt.Println(s)
		}
	}

	// If the command will change data, then enter the aof mode
	if ins.Cmd.WillChangeData && flag != 0 {
		aof.WriteToAof(strings.Join(splits, " ") + "\n")
	}
}
