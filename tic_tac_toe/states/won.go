package states

import "LLD-PRACTICE/tic_tac_toe/models"

type WonState struct{}

func (w *WonState) MakeMove(game *models.Game, row, col int, pType models.PlayerType) error {
	panic("Can't perfom move on won state")
}

func (w *WonState) GetStatus() string {
	return "WON"
}
