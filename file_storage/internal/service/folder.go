package service

import (
	"LLD-PRACTICE/file_storage/internal/interfaces"
)

type Folder struct {
	name     string
	children []interfaces.FileSystem
}

func NewFolder(name string) *Folder {
	return &Folder{
		name:     name,
		children: make([]interfaces.FileSystem, 0),
	}
}

func (f *Folder) GetName() string {
	return f.name
}

func (f *Folder) GetContents() []byte {
	return []byte{}
}

func (f *Folder) AddChildren(child interfaces.FileSystem) {
	f.children = append(f.children, child)
}

func (f *Folder) GetChildNames() []string {
	folderChildrens := []string{f.name}

	for _, child := range f.children {
		folderChildrens = append(folderChildrens, child.GetChildNames()...)
	}

	return folderChildrens
}

func (f *Folder) DeleteFileSystem(name string) {
	for i, children := range f.children {
		if children.GetName() == name {
			// Remove by swapping with last element and truncating it
			f.children[i] = f.children[len(f.children)-1]
			f.children = f.children[:len(f.children)-1]
			return
		}
	}
}
