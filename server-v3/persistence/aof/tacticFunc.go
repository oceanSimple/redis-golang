package aof

import "time"

func alwaysTactic(str string) {
	writeToFile(str)
}

func everySecondTactic() {
	for {
		tempLen := len(InstructionBuf)
		for i := 0; i < tempLen; i++ {
			writeToFile(InstructionBuf[0])
			InstructionBuf = InstructionBuf[1:]
		}
		time.Sleep(time.Second)
	}
}
