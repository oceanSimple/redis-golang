package command

import "server-v2/model"

func listExecute(cmd *model.Command) error {
	switch cmd.Command {
	case "lpush":
		return lpush(cmd)
	case "rpush":
		return rpush(cmd)
	case "lpop":
		return lpop(cmd)
	case "rpop":
		return rpop(cmd)
	case "llen":
		return llen(cmd)
	case "lrange":
		return lrange(cmd)
	case "lindex":
		return lindex(cmd)
	case "linsert":
		return linsert(cmd)
	default:
		return nil
	}
}

func lpush(cmd *model.Command) error {
	return nil
}

func rpush(cmd *model.Command) error {
	return nil
}

func lpop(cmd *model.Command) error {
	return nil
}

func rpop(cmd *model.Command) error {
	return nil
}

func llen(cmd *model.Command) error {
	return nil
}

func lrange(cmd *model.Command) error {
	return nil
}

func lindex(cmd *model.Command) error {
	return nil
}

func linsert(cmd *model.Command) error {
	return nil
}
