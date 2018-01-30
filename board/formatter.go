package board

import (
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

func (board Board) convertRowsToStrings() []string {
	rowStrings := make([]string, board.width)
	rows := board.IndexedRows()

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
