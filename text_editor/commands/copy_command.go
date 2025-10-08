package commands

import (
	"LLD-PRACTICE/text_editor/models"
)

type CopyCommand struct {
	doc    *models.Document
	pos    int
	length int
}

var clipBoard string

func NewCopyCommand(doc *models.Document, pos, length int) *CopyCommand {
	return &CopyCommand{
		doc:    doc,
		pos:    pos,
		length: length,
	}
}

func (c *CopyCommand) Execute() {
	content := c.doc.GetContent()
	if c.pos >= 0 && c.pos+c.length <= len(content) {
		clipBoard = content[c.pos : c.pos+c.length]
	}
}

func (c *CopyCommand) Undo() {} // Copy is non destrective

func (c *CopyCommand) Redo() {
	c.Execute()
}

// copy the content and paste command
