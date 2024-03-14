package instruction

import (
	"fmt"
	"server-v5/persistence/aof"
	"strings"
)

// ExecuteInstruction executes the instruction
// param str: the instruction string
// param flag: the flag of the instruction,
// 0: system --- aof init memory, 1: user --- user input
func ExecuteInstruction(str string, flag int) {
	// 1. Split the string and check the length
	// Now we only support the command which has at least 2 parts
	splits := strings.Fields(str)
	if len(splits) < 2 {
		fmt.Println("ERR wrong number of arguments for 'execute' command")
		return
	}

	// 2. Get the instruction from the ins map
	ins := InsMap[splits[0]]
	if ins == nil { // check if the instruction is nil
		fmt.Println("ERR unknown command '" + splits[0] + "'")
		return
	}

	// 3. Fill key and value
	ins.Key = splits[1]
	// ps: because we have already confirmed the length of the splits is at least 2,
	//     so we can use splits[2:] directly without worrying about the index out of range
	ins.Value = splits[2:]

	// 4. Execute the instruction
	strArr := ins.Execute(ins)

	// 5. Print the result
	if flag != 0 {
		for _, s := range strArr {
			fmt.Println(s)
		}
	}

	// 6. If the command will change data, then enter the aof mode
	if ins.Cmd.WillChangeData && flag != 0 {
		aof.WriteToAof(strings.Join(splits, " ") + "\n")
	}
}
