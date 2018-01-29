package board

const EmptySpace string = "_"

type Board struct {
	spaces []string
}

func (board Board) Spaces() []string {
	return board.spaces
}

func (board Board) NumSpaces() int {
	return len(board.spaces)
}

func (board Board) GetSpace(index int) string {
	return board.spaces[index]
}

func New(dimension int) Board {
	board := Board{}

	board.spaces = make([]string, dimension*dimension)
	for i := range board.spaces {
		board.spaces[i] = EmptySpace
	}
	return board
}

func (board Board) IsMarked(index int) bool {
	return board.spaces[index] != EmptySpace
}

func (board Board) MarkSpace(index int, marker string) {
	board.spaces[index] = marker
}
