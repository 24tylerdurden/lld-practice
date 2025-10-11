package models

type PlayerType string

const (
	X PlayerType = "X"
	O PlayerType = "O"
)

type Player struct {
	ID   string
	Name string
	Type PlayerType
}
