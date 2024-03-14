package aof

import (
	"sync"
)

const (
	// aofFileNane is the path of the aof file
	aofFileNane = "./persistence/aof.aof"
)

var (
	tactic         string   // Tactic is the way to enter the aof mode
	instructionBuf []string // buffer for instructions
)

var (
	instructionBufMutex sync.Mutex = sync.Mutex{} // mutex for instructionBuf
)

var (
	// WriteToAof is the function to write the instruction to the aof file
	WriteToAof func(str string)
	// GoRoutineMethod is the method to run the aof go routine
	GoRoutineMethod func()
)
