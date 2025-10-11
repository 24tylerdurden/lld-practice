package models

type GameState interface {
	MakeMove(game *Game, row, col int, pType PlayerType) error
	GetStatus() string
}

type WinDetector interface {
	CheckWin(board [][]PlayerType, pType PlayerType, size int) bool
	CheckDraw(board [][]PlayerType, size int) bool
}

type GameObserver interface {
	OnMove(row, col int, pType PlayerType)
	OnEnd(status string, winner *Player)
}

// define Interface methods for
// 1. state pattern
// 2. winning strategy ()
// 3. GameObservers

// Game States -
// InProgress
// Won state
// Draw state
