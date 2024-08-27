package insMap

import "server-v5-refactor-server/static/structure"

func PingIns() *structure.Instruction {
	return &structure.Instruction{
		Cmd: structure.Command{
			Value:          "ping",
			BelongTo:       structure.ObjDefault,
			WillChangeData: false,
		},
		Values: nil,
		Execute: func(instruction *structure.Instruction) (response []string) {
			// TODO implement the logic of the instruction
			return []string{"PONG"}
		},
	}
}
