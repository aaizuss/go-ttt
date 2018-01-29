package board

import "errors"

var (
	NoWinnerError = errors.New("This board does not have a winner.")
)

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

func all(spaces []string, f func(string) bool) bool {
	for _, space := range spaces {
		if !f(space) {
			return false
		}
	}
	return true
}

func (board Board) IsFull() bool {
	return all(board.spaces, isMarked)
}

func isWinningRow(row []string) bool {
	firstMarker := row[0]

	return all(row, func(space string) bool {
		return isMarked(space) && space == firstMarker
	})
}

func (board Board) HasWinner() bool {
	for _, row := range board.AllPossibleRowCombos() {
		if isWinningRow(row) {
			return true
		}
	}
	return false
}

func (board Board) Winner() (winner string, err error) {
	for _, row := range board.AllPossibleRowCombos() {
		if isWinningRow(row) {
			return row[0], nil
		}
	}
	return "", NoWinnerError
}

func (board Board) IsTie() bool {
	return board.IsFull() && !board.HasWinner()
}
