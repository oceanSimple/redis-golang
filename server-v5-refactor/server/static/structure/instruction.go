package structure

type Command struct {
	Value          string // instruction name
	BelongTo       string // belong to which data structure：such as sds, list, hash, set, zset
	WillChangeData bool   // whether it will change data
}

type Instruction struct {
	Cmd    Command  // command information
	Values []string // the values of the command
	// the function to execute the command, return the result to print and the error
	Execute func(*Instruction) (response []string)
}
