package commands

import "LLD-PRACTICE/text_editor/models"

type DeleteCommand struct {
	doc     *models.Document
	pos     int
	length  int
	deleted string
}

func NewDeleteCommand(doc *models.Document, pos int, length int) *DeleteCommand {
	return &DeleteCommand{
		doc:    doc,
		pos:    pos,
		length: length,
	}
}

func (d *DeleteCommand) Execute() {
	d.deleted = d.doc.Delete(d.pos, d.length)
}

func (d *DeleteCommand) Undo() {
	d.doc.Insert(d.pos, d.deleted)
}

func (d *DeleteCommand) Redo() {
	d.deleted = d.doc.Delete(d.pos, d.length)
}
