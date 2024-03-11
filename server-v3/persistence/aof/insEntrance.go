package aof

func alwaysEntrance(str string) {
	alwaysTactic(str)
}

func everySecondEntrance(str string) {
	InstructionBuf = append(InstructionBuf, str)
}
