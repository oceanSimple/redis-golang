package global

func addExpireCommand(m map[string]string) {
	m["expire"] = CmdType.Expire
	m["ttl"] = CmdType.Expire
	m["persist"] = CmdType.Expire
	m["setex"] = CmdType.Expire
	m["setnx"] = CmdType.Expire
}

func addSdsCommand(m map[string]string) {
	m["set"] = CmdType.Sds
	m["get"] = CmdType.Sds
	m["strlen"] = CmdType.Sds
	m["del"] = CmdType.Sds
	m["exists"] = CmdType.Sds

	m["mset"] = CmdType.Sds
	m["mget"] = CmdType.Sds

	m["incr"] = CmdType.Sds
	m["decr"] = CmdType.Sds
	m["incrby"] = CmdType.Sds
	m["decrby"] = CmdType.Sds
}

func addListCommand(m map[string]string) {
	m["lpush"] = CmdType.List
	m["rpush"] = CmdType.List

	m["lpop"] = CmdType.List
	m["rpop"] = CmdType.List

	m["llen"] = CmdType.List
	m["lrange"] = CmdType.List
	m["lindex"] = CmdType.List

	m["linsert"] = CmdType.List
}

func addHashCommand(m map[string]string) {
	m["hset"] = CmdType.Hash
	m["hget"] = CmdType.Hash
	m["hdel"] = CmdType.Hash

	m["hmset"] = CmdType.Hash
	m["hmget"] = CmdType.Hash

	m["hlen"] = CmdType.Hash

	m["hgetall"] = CmdType.Hash
}

func addSetCommand(m map[string]string) {
	m["sadd"] = CmdType.Set
	m["srem"] = CmdType.Set
	m["smembers"] = CmdType.Set
	m["sismember"] = CmdType.Set
	m["scard"] = CmdType.Set
}
