package board

const EmptySpace string = "_"

type Board struct {
	spaces    []string
	width     int
	NumSpaces int
}

func (board Board) Spaces() []string {
	return board.spaces
}

func New(width int) Board {
	numSpaces := width * width
	board := Board{NumSpaces: numSpaces, width: width}
	board.spaces = make([]string, numSpaces)

	for i := range board.spaces {
		board.spaces[i] = EmptySpace
	}
	return board
}

func (board Board) MarkSpace(index int, marker string) {
	board.spaces[index] = marker
}

func (board Board) IsValidMove(index int) bool {
	return board.spaceExists(index) && !isMarked(board.spaces[index])
}

func (board Board) EmptySpaces() []int {
	var emptySpaces []int
	for i, space := range board.spaces {
		if !isMarked(space) {
			emptySpaces = append(emptySpaces, i)
		}
	}
	return emptySpaces
}

func (board Board) ResetSpace(index int) {
	board.spaces[index] = EmptySpace
}

func (board Board) spaceExists(index int) bool {
	return index >= 0 && index < board.NumSpaces
}

func isMarked(space string) bool {
	return space != EmptySpace
}
