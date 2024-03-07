package command

import "server-v2/model"

func hashExecute(cmd *model.Command) error {
	switch cmd.Command {
	case "hset":
		return hset(cmd)
	case "hget":
		return hget(cmd)
	case "hdel":
		return hdel(cmd)
	case "hmset":
		return hmset(cmd)
	case "hmget":
		return hmget(cmd)
	case "hlen":
		return hlen(cmd)
	case "hgetall":
		return hgetall(cmd)
	default:
		return nil
	}
}

func hset(cmd *model.Command) error {
	return nil
}

func hget(cmd *model.Command) error {
	return nil
}

func hdel(cmd *model.Command) error {
	return nil
}

func hmset(cmd *model.Command) error {
	return nil
}

func hmget(cmd *model.Command) error {
	return nil
}

func hlen(cmd *model.Command) error {
	return nil
}

func hgetall(cmd *model.Command) error {
	return nil
}
