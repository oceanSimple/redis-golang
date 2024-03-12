package instruction

type commandTypeEnum struct {
	Sds     string
	List    string
	Hash    string
	Set     string
	ZSet    string
	Expire  string
	Unknown string
}

func (e *commandTypeEnum) init() {
	e.Sds = "sds"
	e.List = "list"
	e.Hash = "hash"
	e.Set = "set"
	e.ZSet = "zset"
	e.Expire = "expire"
	e.Unknown = "unknown"
}
