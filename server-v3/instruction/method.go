package instruction

import "strings"

func ExecuteInstruction(str string) {
	splits := strings.Split(str, " ")
	ins := InsMap[splits[0]]
	ins.Key = splits[1]
	ins.Value = splits[2:]

	err := ins.Execute(ins)
	if err != nil {
		return
	}
}
