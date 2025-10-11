package models

type Board struct {
	Size  int
	Cells [][]PlayerType
}

func NewBoard(size int) *Board {
	cells := make([][]PlayerType, size)

	for i := 0; i < size; i++ {
		cells[i] = make([]PlayerType, size)
	}
	return &Board{
		Size:  size,
		Cells: cells,
	}
}

// Method delegates from child to parent
func (b *Board) MoveTo(row, col int, pType PlayerType) bool {
	if row < 0 || row >= b.Size || col < 0 || col >= b.Size || b.Cells[row][col] != "" {
		return false
	}

	b.Cells[row][col] = pType
	return true
}
