package global

type CmdTypeEnum struct {
	Sds     string
	List    string
	Hash    string
	Set     string
	ZSet    string
	Expire  string
	Unknown string
}

func (e *CmdTypeEnum) init() {
	e.Sds = "sds"
	e.List = "list"
	e.Hash = "hash"
	e.Set = "set"
	e.ZSet = "zset"
	e.Expire = "expire"
	e.Unknown = "unknown"
}
