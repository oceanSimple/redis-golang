package aof

import (
	"server-v5/tool"
	"time"
)

func alwaysRoutine() {
}

func everySecondRoutine() {
	var fileConnector, err = tool.GetFileConnector(aofFileNane)
	if err != nil {
		return
	}
	defer tool.CloseFileConnector(fileConnector)
	// Write the instruction in the buffer to the aof file
	for {
		instructionBufMutex.Lock()
		tempLen := len(instructionBuf)
		for i := 0; i < tempLen; i++ {
			tool.WriteToFile(fileConnector, instructionBuf[0])
			instructionBuf = instructionBuf[1:]
		}
		instructionBufMutex.Unlock()
		time.Sleep(time.Second)
	}
}
