package board

func (board Board) AllPossibleRowCombos() [][]string {
	var allRows [][]string

	for _, row := range board.Rows() {
		allRows = append(allRows, row)
	}
	for _, col := range board.Cols() {
		allRows = append(allRows, col)
	}
	for _, diagonal := range board.Diagonals() {
		allRows = append(allRows, diagonal)
	}

	return allRows
}

func (board Board) IsFull() bool {
	for i := range board.spaces {
		if board.IsEmpty(i) {
			return false
		}
	}
	return true
}

func sameInARow(row []string) bool {
	firstMarker := row[0]

	for _, marker := range row {
		if marker == EmptySpace || marker != firstMarker {
			return false
		}
	}
	return true
}

// Only call when you know there is a winner
func (board Board) winningRow() []string {
	var noWinningRow []string
	for _, row := range board.AllPossibleRowCombos() {
		if sameInARow(row) {
			return row
		}
	}
	return noWinningRow
}

// is it bad to write a function that returns 2 values, like
// (bool, winner) where bool is whether the board has a winner and winner is the marker?
// it's more efficient...
func (board Board) HasWinner() bool {
	for _, row := range board.AllPossibleRowCombos() {
		if sameInARow(row) {
			return true
		}
	}
	return false
}

func (board Board) WinningMarker() string {
	winningRow := board.winningRow()
	if len(winningRow) != 0 {
		return winningRow[0]
	} else {
		// this is so ugly
		return ""
	}
}
