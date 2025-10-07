package interfaces

type FileSystem interface {
	GetName() string
	GetContents() []byte
	GetChildNames() []string
}
