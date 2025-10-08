package models

type Command interface {
	Execute()
	Undo()
	Redo()
}

type MacroCommand interface {
	Command
	Add(cmd Command)
}
