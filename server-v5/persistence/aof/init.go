package aof

import (
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"server-v5/log"
)

func init() {
	initTactic()
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

func initInstructionBuf() {
	// If the tactic is "everySecond", buf will be used
	if tactic == "everySecond" {
		instructionBuf = make([]string, 0, 100)
	} else {
		instructionBuf = make([]string, 0)
	}
}

func initInstructionBufWhenRewriteAof() {
	InsBufWhenRewriteAof = make([]string, 50)
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
