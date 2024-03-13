package aof

import "os"

const (
	// aofFileNane is the path of the aof file
	aofFileNane = "./persistence/aof.aof"
)

var (
	tactic         string   // Tactic is the way to enter the aof mode
	fileConnector  *os.File // connection to the aof file
	instructionBuf []string // buffer for instructions
)

var (
	// WriteToAof is the function to write the instruction to the aof file
	WriteToAof func(str string)
	// GoRoutineMethod is the method to run the aof go routine
	GoRoutineMethod func()
)

func GetFileConnector() *os.File {
	return fileConnector
}
