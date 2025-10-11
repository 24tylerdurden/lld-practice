package main

import (
	"LLD-PRACTICE/tic_tac_toe/models"
	"LLD-PRACTICE/tic_tac_toe/services"
	"fmt"
)

func main() {
	factory := services.NewGameFactory()

	player1 := &models.Player{ID: "p1", Name: "Pavan", Type: models.X}
	player2 := models.Player{ID: "p2", Name: "Teja", Type: models.O}

	game := factory.Create3x3Board(player1, &player2)

	// PlayMoves
	game.MakeMove(0, 0) // X
	game.MakeMove(1, 1) // O
	game.MakeMove(0, 1) // X
	game.MakeMove(1, 0) // O
	game.MakeMove(0, 2) // X → Wins!

	fmt.Println("The board state is : ", game.Board)

}
