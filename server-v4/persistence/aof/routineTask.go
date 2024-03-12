package aof

import "time"

func alwaysRoutine() {
}

func everySecondRoutine() {
	for {
		tempLen := len(instructionBuf)
		for i := 0; i < tempLen; i++ {
			writeToFile(instructionBuf[0])
			instructionBuf = instructionBuf[1:]
		}
		time.Sleep(time.Second)
	}
}
