package board

import (
	"strconv"
	"strings"
)

func (board Board) FormattedString() string {
	divider := divider()
	stringRows := board.convertRowsToStrings()
	firstRow := stringRows[0]
	secondRow := stringRows[1]
	thirdRow := stringRows[2]

	return firstRow + divider + secondRow + divider + thirdRow
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
	result := " " + strings.Join(row, " | ") + "\n"
	return result
}

func divider() string {
	return strings.Repeat("-", 11) + "\n"
}

func (board Board) indexedRows() [][]string {
	newSpaces := make([]string, board.NumSpaces)
	for i, space := range board.spaces {
		newSpaces[i] = renderSpace(i, space)
	}
	indexedBoard := Board{spaces: newSpaces, NumSpaces: board.NumSpaces, width: board.width}
	return indexedBoard.Rows()
}
