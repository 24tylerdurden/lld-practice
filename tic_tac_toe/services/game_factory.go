package services

import (
	"LLD-PRACTICE/tic_tac_toe/models"
	"LLD-PRACTICE/tic_tac_toe/observers"
	"LLD-PRACTICE/tic_tac_toe/states"
	"LLD-PRACTICE/tic_tac_toe/strategies"
)

type GameFactory struct {
}

func NewGameFactory() *GameFactory {
	return &GameFactory{}
}

func (g *GameFactory) Create3x3Board(player1, player2 *models.Player) *models.Game {
	obs := []models.GameObserver{&observers.GameObserver{}}
	game := models.NewGame(player1, player2, 3)
	game.SetState(&states.InProgress{})
	game.AddObservers(obs)
	game.SetStrategy(&strategies.StandardWinDetector{})

	return game
}

func (g *GameFactory) CreateNXNBoard(player1, player2 *models.Player) *models.Game {

	obs := []models.GameObserver{}
	game := models.NewGame(player1, player2, 5)
	game.SetState(&states.InProgress{})
	game.AddObservers(obs)

	return game
}
