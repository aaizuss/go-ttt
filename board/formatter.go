package board

import (
	"strconv"
	"strings"
)

const (
	blank        = " "
	spaceDivider = " | "
	rowLine      = "\n-----------\n"
)

func (board Board) FormattedString() string {
	stringRows := board.convertRowsToStrings()
	firstRow := stringRows[0]
	secondRow := stringRows[1]
	thirdRow := stringRows[2]

	return firstRow + rowLine + secondRow + rowLine + thirdRow + "\n"
}

func renderSpace(index int, space string) string {
	if isMarked(space) {
		return space
	} else {
		return strconv.Itoa(index)
	}
}

func (board Board) convertRowsToStrings() []string {
	rowStrings := make([]string, board.width)
	rows := board.indexedRows()

	for i, row := range rows {
		rowStrings[i] = rowToString(row)
	}
	return rowStrings
}

func rowToString(row []string) string {
	result := blank + strings.Join(row, spaceDivider)
	return result
}

func (board Board) indexedRows() [][]string {
	newSpaces := make([]string, board.NumSpaces)
	for i, space := range board.spaces {
		newSpaces[i] = renderSpace(i, space)
	}
	indexedBoard := Board{spaces: newSpaces, NumSpaces: board.NumSpaces, width: board.width}
	return indexedBoard.Rows()
}
