package aof

import (
	"os"
	"server-v5-refactor-server/persistence"
	"server-v5-refactor-server/setting"
	"server-v5-refactor-server/static/structure"
	"strings"
	"sync"
)

var (
	aofFileConnector *os.File
	bufferBuilder    strings.Builder
	tactic           = "always"
)

var (
	connMutex   = sync.Mutex{}
	bufferMutex = sync.Mutex{}
)

const (
	always      = "always"
	everySecond = "everySecond"
	no          = "no"
)

func init() {
	aofFileConnector, _ = persistence.GetFileConnector("aof.txt")
	bufferBuilder = strings.Builder{}
	tactic = setting.GetDefault("aof.tactic").(string)
}

func WriteToAof(ins *structure.Instruction) {
	// judge the instruction whether you need to be written to aof
	if !ins.Cmd.WillChangeData {
		return
	}
	msg := spliceInstruction(ins)
	if tactic == always { // write to aof file immediately
		connMutex.Lock()
		persistence.WriteToFile(aofFileConnector, msg)
		connMutex.Unlock()
	} else { // write to buffer
		bufferMutex.Lock()
		bufferBuilder.WriteString(msg)
		bufferMutex.Unlock()
	}
}

func FlushBuffer() {
	bufferMutex.Lock()
	connMutex.Lock()
	persistence.WriteToFile(aofFileConnector, bufferBuilder.String())
	bufferBuilder.Reset()
	connMutex.Unlock()
	bufferMutex.Unlock()
}

func GetTactic() string {
	return tactic
}

func spliceInstruction(ins *structure.Instruction) string {
	return ins.Cmd.Value + " " + strings.Join(ins.Values, " ") + "\n"
}
