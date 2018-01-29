package board

const EmptySpace string = "_"

func New(size int) []string {
	board := make([]string, size*size)
	for i := range board {
		board[i] = EmptySpace
	}
	return board
}

func MarkSpace(board []string, space int, marker string) []string {
	board[space] = marker
	return board
}
