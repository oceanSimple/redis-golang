package instruction

type CommandTypeEnum struct {
	Sds     string
	List    string
	Hash    string
	Set     string
	ZSet    string
	Expire  string
	Unknown string
}

func (e *CommandTypeEnum) init() {
	e.Sds = "sds"
	e.List = "list"
	e.Hash = "hash"
	e.Set = "set"
	e.ZSet = "zset"
	e.Expire = "expire"
	e.Unknown = "unknown"
}
