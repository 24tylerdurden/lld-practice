package states

import (
	"LLD-PRACTICE/tic_tac_toe/models"
	"fmt"
)

type InProgress struct{}

func (i *InProgress) MakeMove(game *models.Game, row, col int, pType models.PlayerType) error {

	fmt.Println("This is being set called and set here")

	// Execute Move, If It's a valid state we will be mutating the Game state here
	if !game.Board.MoveTo(row, col, pType) {
		return fmt.Errorf("Invalid Move")
	}

	// Notify Observers
	for _, obs := range game.Observers {
		obs.OnMove(row, col, pType)
	}

	// Check Win
	if game.ResDetector.CheckWin(game.Board.Cells, pType, game.Board.Size) {
		game.Winner = game.GetCurrentPlayer()
		game.State = &WonState{}

		for _, obs := range game.Observers {
			fmt.Println("")
			obs.OnEnd("WON", game.Winner)
		}

		return nil
	}

	// Check Draw

	if game.ResDetector.CheckDraw(game.Board.Cells, game.Board.Size) {

		game.State = &DrawState{}

		for _, obs := range game.Observers {
			obs.OnEnd("DRAW", nil)
		}
		return nil
	}

	// Swicth Player
	game.SwitchPlayer()
	return nil
}

func (i *InProgress) GetStatus() string {
	return "IN_PROGRESS"
}
