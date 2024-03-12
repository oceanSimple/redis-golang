package model

type Command struct {
	Value          string // instruction name
	BelongTo       string // belong to which data structure：such as sds, list, hash, set, zset
	WillChangeData bool   // whether it will change data
}

type Instruction struct {
	Cmd   Command  // command information
	Key   string   // the key of the command
	Value []string // the value of the command
	// the function to execute the command, return the result to print and the error
	Execute func(*Instruction) ([]string, error)
}
