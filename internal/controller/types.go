package controller

type Command string

const (
	Add     Command = "add"
	List    Command = "list"
	Delete  Command = "delete"
	Summary Command = "summary"
	Update  Command = "update"
)
