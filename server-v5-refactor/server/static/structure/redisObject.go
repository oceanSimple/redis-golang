package structure

type RedisObject struct {
	Type string      // Type of the object
	Ptr  interface{} // value of the object
}
