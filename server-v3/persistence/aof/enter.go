package aof

import (
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"os"
	"server-v3/log"
)

var (
	Tactic          string // Tactic is the way to enter the aof mode
	FileConnector   *os.File
	InstructionBuf  []string
	InsEntrance     func(str string)
	GoRoutineMethod func()
)

var (
	aofFileNane = "./persistence/aof.aof"
)

func init() {
	tacticInit()
	fileConnectorInit()
	instructionBufInit()
	insEntranceInit()
	goRoutineMethodInit()
}

func tacticInit() {
	Tactic = viper.GetString("aof.tactic")
}

func fileConnectorInit() {
	var err error
	FileConnector, err = os.OpenFile(aofFileNane,
		os.O_CREATE|os.O_APPEND|os.O_RDWR,
		0666)
	if err != nil {
		log.SystemLog.Fatal("aof file open failed",
			zap.Error(err))
	}
}

func instructionBufInit() {
	InstructionBuf = make([]string, 0, 100)
}

func insEntranceInit() {
	switch Tactic {
	case "always":
		InsEntrance = alwaysEntrance
	case "everySecond":
		InsEntrance = everySecondEntrance
	default:
		InsEntrance = alwaysEntrance
	}
}

func goRoutineMethodInit() {
	var emptyFunc = func() {}
	switch Tactic {
	case "always":
		GoRoutineMethod = emptyFunc
	case "everySecond":
		GoRoutineMethod = everySecondTactic
	default:
		GoRoutineMethod = emptyFunc
	}
}
