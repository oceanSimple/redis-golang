package enum

type typeEnum struct {
	String string
	List   string
	Hash   string
	Set    string
	ZSet   string
}

func (t typeEnum) init() {
	t.ZSet = "zset"
	t.Set = "set"
	t.Hash = "hash"
	t.List = "list"
	t.String = "string"
}
