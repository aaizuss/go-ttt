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
	return board.spaceExists(index) && board.spaceEmpty(index)
}

func (board Board) EmptySpaces() []int {
	var emptySpaces []int
	for index := range board.spaces {
		if board.spaceEmpty(index) {
			emptySpaces = append(emptySpaces, index)
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

func (board Board) spaceEmpty(index int) bool {
	return board.spaces[index] == EmptySpace
}
