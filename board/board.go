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

// currently only used in tests - might delete
func (board Board) GetSpace(index int) string {
	return board.spaces[index]
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

// currently only used in tests - might delete
func (board Board) IsMarked(index int) bool {
	return board.spaces[index] != EmptySpace
}

// might be able to delete this, just use isMarked (but I think I use IsEmpty in tests)
func (board Board) IsEmpty(index int) bool {
	return board.spaces[index] == EmptySpace
}

func isMarked(space string) bool {
	return space != EmptySpace
}

func (board Board) MarkSpace(index int, marker string) {
	board.spaces[index] = marker
}

func (board Board) SpaceExists(index int) bool {
	return index >= 0 && index < board.NumSpaces
}
