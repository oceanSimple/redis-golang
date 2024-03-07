package command

import (
	"server-v2/global"
	"server-v2/model"
	"strings"
)

type CmdParseInterface interface {
	ExecuteCmd(str string) error
	ParseCmdStr(str string) *model.Command
	JudgeCmdType(cmd *model.Command) string
	SelectCmdExecuteFunc(cmdType string, cmd *model.Command) error
}

type CmdParse struct {
}

func (c *CmdParse) ExecuteCmd(str string) error {
	cmd := c.ParseCmdStr(str)
	cmdType := c.JudgeCmdType(cmd)
	return c.SelectCmdExecuteFunc(cmdType, cmd)
}

func (c *CmdParse) ParseCmdStr(str string) *model.Command {
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

func (c *CmdParse) JudgeCmdType(cmd *model.Command) string {
	if value, ok := global.CmdMap[cmd.Command]; ok {
		return value
	} else {
		return global.CmdType.Unknown
	}
}

func (c *CmdParse) SelectCmdExecuteFunc(cmdType string, cmd *model.Command) error {
	switch cmdType {
	case global.CmdType.Sds:
		return sdsExecute(cmd)
	case global.CmdType.List:
		return listExecute(cmd)
	case global.CmdType.Hash:
		return hashExecute(cmd)
	case global.CmdType.Set:
		return setExecute(cmd)
	case global.CmdType.ZSet:
		return zsetExecute(cmd)
	case global.CmdType.Expire:
		return expireExecute(cmd)
	default:
		return unknownExecute(cmd)
	}
}
