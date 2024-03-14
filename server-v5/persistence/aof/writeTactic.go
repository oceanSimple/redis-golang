package aof

import "server-v5/tool"

func alwaysTactic(str string) {
	var fileConnector, err = tool.GetFileConnector(aofFileNane)
	if err != nil {
		return
	}
	defer tool.CloseFileConnector(fileConnector)
	// Write the instruction to the aof file
	tool.WriteToFile(fileConnector, str)
}

func everySecondTactic(str string) {
	// Write the instruction to the buffer
	instructionBufMutex.Lock()
	instructionBuf = append(instructionBuf, str)
	instructionBufMutex.Unlock()
}
