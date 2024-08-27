package insMap

import "server-v5-refactor-server/static/structure"

var InstructionMap map[string]*structure.Instruction

func init() {
	InstructionMap = make(map[string]*structure.Instruction)
	defaultIns(InstructionMap)
	strIns(InstructionMap)
}

func defaultIns(m map[string]*structure.Instruction) {
	m["ping"] = PingIns()
}

func strIns(m map[string]*structure.Instruction) {
	m["get"] = get()
	m["set"] = set()
	m["del"] = del()
	m["strlen"] = strlen()
	m["append"] = appendStr()
	m["incr"] = incr()
	m["decr"] = decr()
	m["incrby"] = incrby()
	m["decrby"] = decrby()
}
