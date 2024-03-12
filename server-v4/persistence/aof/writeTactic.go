package aof

func alwaysTactic(str string) {
	writeToFile(str)
}

func everySecondTactic(str string) {
	instructionBuf = append(instructionBuf, str)
}
