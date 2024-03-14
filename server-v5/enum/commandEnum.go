package enum

type commandTypeEnum struct {
	String  string
	List    string
	Hash    string
	Set     string
	ZSet    string
	Expire  string
	Unknown string
}

func (e *commandTypeEnum) init() {
	e.String = "string"
	e.List = "list"
	e.Hash = "hash"
	e.Set = "set"
	e.ZSet = "zset"
	e.Expire = "expire"
	e.Unknown = "unknown"
}
