package tool

import (
	"fmt"
	"server-v2/global"
	"server-v2/model"
	"strings"
)

type commandToolInterface interface {
	ParseCommand(str string) *model.Command
	ExecuteCommand(cmd model.Command) error
}

type commandTool struct {
}

func (c commandTool) ParseCommand(str string) *model.Command {
	strSplit := strings.Split(str, " ")
	switch len(strSplit) {
	case 0:
		return &model.Command{}
	case 1:
		return &model.Command{
			Command: strSplit[0],
		}
	case 2:
		return &model.Command{
			Command: strSplit[0],
			Key:     strSplit[1],
		}
	default:
		return &model.Command{
			Command: strSplit[0],
			Key:     strSplit[1],
			Value:   strSplit[2:],
		}
	}
}

func (c commandTool) ExecuteCommand(cmd model.Command) error {
	//TODO implement me
	//just implement the sds command
	switch cmd.Command {
	case "set":
		global.Map.SdsMap[cmd.Key] = model.Sds(strings.Join(cmd.Value, " "))
		return nil
	case "get":
		if v, ok := global.Map.SdsMap[cmd.Key]; ok {
			fmt.Println(">", v)
		}
		return nil
	default:
		fmt.Println("command not found")
		return fmt.Errorf("command not found")
	}
}
