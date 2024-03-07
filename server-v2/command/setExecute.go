package command

import "server-v2/model"

func setExecute(cmd *model.Command) error {
	switch cmd.Command {
	case "sadd":
		return sadd(cmd)
	case "srem":
		return srem(cmd)
	case "smembers":
		return smembers(cmd)
	case "sismember":
		return sismember(cmd)
	case "scard":
		return scard(cmd)
	default:
		return nil
	}
}

func sadd(cmd *model.Command) error {
	return nil
}

func srem(cmd *model.Command) error {
	return nil
}

func smembers(cmd *model.Command) error {
	return nil
}

func sismember(cmd *model.Command) error {
	return nil
}

func scard(cmd *model.Command) error {
	return nil
}
