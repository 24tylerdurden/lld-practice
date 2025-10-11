package strategies

import "LLD-PRACTICE/tic_tac_toe/models"

type StandardWinDetector struct {
}

func (s *StandardWinDetector) CheckWin(board [][]models.PlayerType, pType models.PlayerType, size int) bool {

	if size != 3 {
		return s.fallbackCheck(board, pType, size)
	}

	// Check Row
	for i := 0; i < 3; i++ {
		if board[i][0] == pType && board[i][1] == pType && board[i][2] == pType {
			return true
		}
	}

	// Check Col
	for i := 0; i < 3; i++ {
		if board[0][i] == pType && board[1][i] == pType && board[2][i] == pType {
			return true
		}
	}

	// Check Diagonals

	if board[0][0] == pType && board[1][1] == pType && board[2][2] == pType {
		return true
	}

	if board[0][2] == pType && board[1][1] == pType && board[2][0] == pType {
		return true
	}

	return false
}

func (s *StandardWinDetector) CheckDraw(board [][]models.PlayerType, size int) bool {

	//  check if there are any empty cells

	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			if board[i][j] == "" {
				return false
			}
		}
	}

	return true
}

func (s *StandardWinDetector) fallbackCheck(board [][]models.PlayerType, pType models.PlayerType, size int) bool {

	scalable := &ScalableWinDetector{}

	return scalable.CheckWin(board, pType, size)

}

// Check Win
// Check Draw
