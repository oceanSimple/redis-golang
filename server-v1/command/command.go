package command

import (
	"fmt"
	"server-v1/global"
	"strings"
)

type CmdInterface interface {
	ParseCommand(cmdStr string) *Command
	ExecuteCommand(cmd *Command) error
}

type Command struct {
	Command string
	Key     string
	Value   []string
}

func (c *Command) ExecuteCommand(cmd *Command) error {
	var err error = nil
	switch cmd.Command {
	case "set":
		global.GlobalMap.SdsMap[cmd.Key] =
			global.GlobalBuilder.SdsBuilder.NewSds(cmd.Value[0])
	case "get":
		val := global.GlobalMap.SdsMap[cmd.Key].String()
		fmt.Println(val)
	case "list":
		for k, v := range global.GlobalMap.SdsMap {
			fmt.Println(k, v)
		}
	default:
		fmt.Println("invalid command")
		err = fmt.Errorf("invalid command")
	}
	return err
}

func (c *Command) ParseCommand(cmdStr string) *Command {
	cmdSplit := strings.Split(cmdStr, " ")
	if len(cmdSplit) < 1 {
		return nil
	}
	if len(cmdSplit) == 1 {
		return &Command{
			Command: cmdSplit[0],
		}
	}
	if len(cmdSplit) == 2 {
		return &Command{
			Command: cmdSplit[0],
			Key:     cmdSplit[1],
		}
	}
	if len(cmdSplit) > 2 {
		return &Command{
			Command: cmdSplit[0],
			Key:     cmdSplit[1],
			Value:   cmdSplit[2:],
		}
	}
	return nil
}
