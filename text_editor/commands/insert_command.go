package commands

import "LLD-PRACTICE/text_editor/models"

type InsertCommand struct {
	doc  *models.Document
	pos  int
	text string
}

func NewInsertCommand(doc *models.Document, pos int, text string) *InsertCommand {
	return &InsertCommand{
		pos:  pos,
		text: text,
		doc:  doc,
	}
}

// Implements Cmd Interaface

func (I *InsertCommand) Execute() {
	I.doc.Insert(I.pos, I.text)
}

func (I *InsertCommand) Undo() {
	I.doc.Delete(I.pos, len(I.text))
}

func (I *InsertCommand) Redo() {
	I.Execute()
}
