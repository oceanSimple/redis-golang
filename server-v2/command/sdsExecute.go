package command

import (
	"fmt"
	"server-v2/global"
	"server-v2/model"
)

func sdsExecute(cmd *model.Command) error {
	switch cmd.Command {
	case "set":
		return set(cmd)
	case "get":
		return get(cmd)
	case "strlen":
		return strlen(cmd)
	case "del":
		return del(cmd)
	case "mset":
		return mset(cmd)
	case "mget":
		return mget(cmd)
	case "incr":
		return incr(cmd)
	case "decr":
		return decr(cmd)
	case "incrby":
		return incrby(cmd)
	case "decrby":
		return decrby(cmd)
	default:
		return nil
	}
}

func set(cmd *model.Command) error {
	global.Map.SdsMap[cmd.Key] = model.Sds(cmd.Value[0])
	return nil
}

func get(cmd *model.Command) error {
	fmt.Println(global.Map.SdsMap[cmd.Key])
	return nil
}

func strlen(cmd *model.Command) error {
	return nil
}

func del(cmd *model.Command) error {
	return nil
}

func mset(cmd *model.Command) error {
	return nil
}

func mget(cmd *model.Command) error {
	return nil
}

func incr(cmd *model.Command) error {
	return nil
}

func decr(cmd *model.Command) error {
	return nil
}

func incrby(cmd *model.Command) error {
	return nil
}

func decrby(cmd *model.Command) error {
	return nil
}

//func exists(cmd *model.Command) error {
//	return nil
//}
