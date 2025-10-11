package observers

import (
	"LLD-PRACTICE/tic_tac_toe/models"
	"fmt"
)

type GameObserver struct {
}

func (g *GameObserver) OnMove(row, col int, pType models.PlayerType) {
	fmt.Printf("Move: Player %s at (%d, %d)\n", pType, row, col)
}

func (g *GameObserver) OnEnd(status string, winner *models.Player) {
	if status == "WON" {
		fmt.Printf(" Winner: %s\n", winner.Name)
	} else {
		fmt.Printf(" It's a draw")
	}
}
