package states

import "LLD-PRACTICE/tic_tac_toe/models"

type DrawState struct{}

func (d *DrawState) MakeMove(game *models.Game, row, col int, pType models.PlayerType) error {
	panic("Can't perform move on Draw state")
}

func (d *DrawState) GetStatus() string {
	return "DRAW"
}
