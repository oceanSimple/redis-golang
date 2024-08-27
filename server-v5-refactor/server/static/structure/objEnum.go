package structure

type ObjType string

const (
	ObjString  = "STRING"
	ObjList    = "LIST"
	ObjSet     = "SET"
	ObjHash    = "HASH"
	ObjZset    = "ZSET"
	ObjDefault = "DEFAULT"
)
