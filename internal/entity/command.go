package entity

type Command map[string]interface{}

type CommandResult interface{}

const (
	CmdSetTxId string = "SetTxId"
)
