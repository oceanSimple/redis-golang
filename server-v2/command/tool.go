package command

import (
	"fmt"
	"server-v2/global"
	"server-v2/model"
	"strings"
)

type ToolInterface interface {
	ParseCommand(cmdStr string) *model.Command
	ExecuteCommand(cmd *model.Command) error
}

type Tool struct {
}

func (c *Tool) ExecuteCommand(cmd *model.Command) error {
	var err error = nil
	fmt.Println("打印cmd: ")
	fmt.Println(cmd.Command)
	switch cmd.Command {
	case "set":
		if cmd.Key == "" || len(cmd.Value) == 0 {
			return fmt.Errorf("invalid key or value")
		}
		global.Map.SdsMap[cmd.Key] = model.Sds(cmd.Value[0])
	case "get":
		fmt.Println("get key: ", cmd.Key)
		val, ok := global.Map.SdsMap[cmd.Key]
		if !ok {
			fmt.Println("not found")
		} else {
			fmt.Println(val)
		}
	case "list":
		if len(global.Map.SdsMap) > 1000 { // replace 1000 with the appropriate size limit
			return fmt.Errorf("too many items to list")
		}
		for k, v := range global.Map.SdsMap {
			fmt.Println(k, v)
		}
	default:
		fmt.Println("invalid command")
		err = fmt.Errorf("invalid command")
	}
	return err
}

func (c *Tool) ParseCommand(cmdStr string) *model.Command {
	cmdStr = strings.Trim(cmdStr, "\n")
	fmt.Println("#cmdStr: ", cmdStr)
	cmdSplit := strings.Split(cmdStr, " ")
	fmt.Println("打印cmdSplit: ")
	for i, s := range cmdSplit {
		fmt.Println(i, s)
	}
	if len(cmdSplit) < 1 {
		return nil
	}
	if len(cmdSplit) == 1 {
		return &model.Command{
			Command: cmdSplit[0],
			Key:     "",
			Value:   []string{""},
		}
	}
	if len(cmdSplit) == 2 {
		return &model.Command{
			Command: cmdSplit[0],
			Key:     cmdSplit[1],
			Value:   []string{""},
		}
	}
	if len(cmdSplit) > 2 {
		return &model.Command{
			Command: cmdSplit[0],
			Key:     cmdSplit[1],
			Value:   cmdSplit[2:],
		}
	}
	return nil
}
