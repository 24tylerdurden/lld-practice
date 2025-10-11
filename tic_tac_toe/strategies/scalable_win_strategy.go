package strategies

import "LLD-PRACTICE/tic_tac_toe/models"

type ScalableWinDetector struct{}

func (s *ScalableWinDetector) CheckWin(cells [][]models.PlayerType, pType models.PlayerType, size int) bool {
	// Check rows

	for i := 0; i < size; i++ {
		win := true

		for j := 0; j < size; j++ {
			if cells[i][j] != pType {
				win = false
				break
			}
		}

		if win {
			return true
		}
	}

	// check columns

	for i := 0; i < size; i++ {
		win := true

		for j := 0; j < size; j++ {
			if cells[i][j] != pType {
				win = false
				break
			}
		}

		if win {
			return true
		}

	}

	// check diagonals main diagonal
	win := true

	for i := 0; i < size; i++ {
		if cells[i][i] != pType {
			win = false
			break
		}
	}

	if win {
		return true
	}

	// check anti diagonal

	win = true

	for i := 0; i < size; i++ {
		if cells[i][size-i-1] != pType {
			win = false
			break
		}
	}

	if win {
		return win
	}

	return false

}

func (s *ScalableWinDetector) CheckDraw(cells [][]models.PlayerType, size int) bool {

	// check for any empty cells

	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			if cells[i][j] == "" {
				return false
			}
		}
	}

	return true

}
