package entity

type CommandContext string

type Command map[CommandContext]interface{}

type CommandResult map[string]interface{}
