package aof

import (
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"os"
	"server-v4/log"
)

func init() {
	initTactic()
	initFileConnector()
	initInstructionBuf()
	initFunc()
}

func initTactic() {
	tactic = viper.GetString("aof.tactic")
	// If the tactic is not "always" or "everySecond", set it to "always"
	if tactic != "always" && tactic != "everySecond" {
		tactic = "always"
	}
}

func initFileConnector() {
	var err error
	fileConnector, err = os.OpenFile(aofFileNane,
		os.O_CREATE|os.O_APPEND|os.O_RDWR,
		0666)
	if err != nil {
		log.SystemLog.Fatal("aof file open failed",
			zap.Error(err))
	}
}

func initInstructionBuf() {
	// If the tactic is "everySecond", buf will be used
	if tactic == "everySecond" {
		instructionBuf = make([]string, 0, 100)
	} else {
		instructionBuf = make([]string, 0)
	}
}

func initFunc() {
	if tactic == "always" {
		WriteToAof = alwaysTactic
		GoRoutineMethod = alwaysRoutine
	} else if tactic == "everySecond" {
		WriteToAof = everySecondTactic
		GoRoutineMethod = everySecondRoutine
	} else {
		log.SystemLog.Fatal("aof tactic error",
			zap.String("tactic", tactic))
	}
}
