package main

import (
	"LLD-PRACTICE/text_editor/commands"
	"LLD-PRACTICE/text_editor/models"
	"LLD-PRACTICE/text_editor/services"
	"fmt"
)

func main() {
	doc := models.NewDocument()
	history := services.NewHistoryManager(100)

	// Insert "Hello"
	insertCmd := commands.NewInsertCommand(doc, 0, "Hello")
	history.ExecuteCommand(insertCmd)
	fmt.Println("After insert:", doc.GetContent())

	// Delete "ello"
	deleteCmd := commands.NewDeleteCommand(doc, 1, 4)
	history.ExecuteCommand(deleteCmd)
	fmt.Println("After delete:", doc.GetContent())

	// Undo delete
	history.Undo()
	fmt.Println("After undo:", doc.GetContent())

	// Redo delete
	history.Redo()
	fmt.Println("After redo:", doc.GetContent())

	// Macro: Insert " World!" + Copy + Paste
	macro := commands.NewMacroCommand()
	macro.Add(commands.NewInsertCommand(doc, 5, " World!"))
	macro.Add(commands.NewCopyCommand(doc, 6, 5)) // "World"

	history.ExecuteCommand(macro)
	fmt.Println("After macro:", doc.GetContent())
}
