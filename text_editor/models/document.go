package models

import "sync"

type Document struct {
	content string
	mu      sync.Mutex
}

func NewDocument() *Document {
	return &Document{
		content: "",
	}
}

// Insert, delete, getContent, getLength

func (d *Document) Insert(pos int, cont string) {
	d.mu.Lock()
	defer d.mu.Unlock()

	if pos < 0 {
		return
	}

	d.content = d.content + cont
}

func (d *Document) Delete(pos, length int) string {
	d.mu.Lock()
	defer d.mu.Unlock()

	if pos < 0 || pos+length > len(d.content) {
		return ""
	}

	deleted := d.content[pos : pos+length]
	d.content = d.content[:pos] + d.content[pos+length:]
	return deleted
}

func (d *Document) GetContent() string {
	return d.content
}

func (d *Document) GetLen() int {
	return len(d.content)
}
