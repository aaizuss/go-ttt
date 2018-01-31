package board

import "errors"

var (
	NoWinnerError = errors.New("this board does not have a winner")
)

func (board Board) WinningCombos() [][]string {
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

func (board Board) Winner() (winner string, err error) {
	for _, row := range board.WinningCombos() {
		if isWinningRow(row) {
			return row[0], nil
		}
	}
	return "", NoWinnerError
}

func (board Board) HasWinner() bool {
	for _, row := range board.WinningCombos() {
		if isWinningRow(row) {
			return true
		}
	}
	return false
}

func (board Board) IsFull() bool {
	return all(board.spaces, isMarked)
}

func (board Board) IsTie() bool {
	return board.IsFull() && !board.HasWinner()
}

func (board Board) GameOver() bool {
	return board.IsTie() || board.HasWinner()
}

func isWinningRow(row []string) bool {
	firstMarker := row[0]

	return all(row, func(space string) bool {
		return isMarked(space) && space == firstMarker
	})
}

func all(spaces []string, f func(string) bool) bool {
	for _, space := range spaces {
		if !f(space) {
			return false
		}
	}
	return true
}
