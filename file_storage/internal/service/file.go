package service

import "sync"

type File struct {
	name    string
	content []byte
	mu      sync.RWMutex
}

func NewFile(name string, content []byte) *File {
	return &File{
		name:    name,
		content: content,
	}
}

func (f *File) GetName() string {
	return f.name
}

func (f *File) GetContents() []byte {
	return f.content
}

func (f *File) GetChildNames() []string {
	return []string{f.name}
}
