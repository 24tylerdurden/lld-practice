package models

import (
	"fmt"
	"sync"
)

type Game struct {
	Board         *Board
	ResDetector   WinDetector
	State         GameState
	mu            sync.Mutex
	currentPlayer PlayerType
	Observers     []GameObserver
	players       [2]*Player
	Winner        *Player
}

func NewGame(player1, player2 *Player, size int) *Game {
	return &Game{
		players:       [2]*Player{player1, player2},
		Board:         NewBoard(size),
		currentPlayer: "X",
	}
}

func (g *Game) MakeMove(row, col int) {

	g.mu.Lock()
	defer g.mu.Unlock()

	fmt.Println("cc1")

	g.State.MakeMove(g, row, col, g.currentPlayer)
}

func (g *Game) SwitchPlayer() {
	fmt.Println("Inside the switch state method")
	if g.currentPlayer == "X" {
		g.currentPlayer = "O"
	} else {
		g.currentPlayer = "X"
	}
}

func (g *Game) GetCurrentPlayer() *Player {
	if g.currentPlayer == "X" {
		return g.players[0]
	}
	return g.players[1]
}

func (g *Game) SetState(state GameState) {
	g.State = state
}

func (g *Game) AddObservers(obs []GameObserver) {
	g.Observers = append(g.Observers, obs...)
}

func (g *Game) SetStrategy(strat WinDetector) {
	g.ResDetector = strat
}
