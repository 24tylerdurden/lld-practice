package commands

import "LLD-PRACTICE/text_editor/models"

// Composite design pattern
type MacroCommands struct {
	commands []models.Command
}

func NewMacroCommand() *MacroCommands {
	return &MacroCommands{}
}

func (m *MacroCommands) Add(cmd models.Command) {
	m.commands = append(m.commands, cmd)
}

func (m *MacroCommands) Execute() {
	for _, cmd := range m.commands {
		cmd.Execute()
	}
}

func (m *MacroCommands) Undo() {
	for i := len(m.commands) - 1; i >= 0; i-- {
		m.commands[i].Undo()
	}
}

func (m *MacroCommands) Redo() {
	m.Execute()
}
