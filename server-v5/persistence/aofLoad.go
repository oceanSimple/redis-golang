package persistence

import (
	"bufio"
	"fmt"
	"go.uber.org/zap"
	"server-v5/instruction"
	"server-v5/log"
	"server-v5/tool"
)

func LoadAofFile() {
	fileConnector, err := tool.GetFileConnector("./persistence/aof.aof")
	if err != nil {
		return
	}
	defer tool.CloseFileConnector(fileConnector)

	//var reader = bufio.NewReader(fileConnector)
	//for {
	//	line, err := reader.ReadString('\n')
	//	if err != nil {
	//		if err == io.EOF {
	//			break
	//		} else {
	//			log.SystemLog.Error("Read aof file error",
	//				zap.Error(err))
	//			return
	//		}
	//	}
	//	fmt.Println(line)
	//	instruction.ExecuteInstruction(line, 0)
	//}

	scanner := bufio.NewScanner(fileConnector)
	for scanner.Scan() {
		instruction.ExecuteInstruction(scanner.Text(), 0)
	}
	if err := scanner.Err(); err != nil {
		fmt.Println("aof file read failed")
		log.SystemLog.Fatal("aof file read failed",
			zap.Error(err))
	}
}
