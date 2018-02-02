package board

func (board Board) Rows() [][]string {
	width := board.width
	rows := make([][]string, width)

	for i := 0; i < width; i++ {
		rowStartIndex := i * width
		rowEndIndex := rowStartIndex + board.width
		rows[i] = board.spaces[rowStartIndex:rowEndIndex]
	}
	return rows
}

func (board Board) Cols() [][]string {
	width := board.width
	rows := board.Rows()
	cols := make([][]string, width)

	for i := 0; i < width; i++ {
		cols[i] = extractCol(rows, i)
	}
	return cols
}

func (board Board) Diagonals() [][]string {
	rows := board.Rows()
	diagonals := make([][]string, 2)

	for i, row := range rows {
		diagonals[0] = append(diagonals[0], row[i])
		diagonals[1] = append(diagonals[1], row[board.width-1-i])
	}

	return diagonals
}

func extractCol(rows [][]string, colIndex int) []string {
	col := make([]string, 3)

	for i, row := range rows {
		col[i] = row[colIndex]
	}
	return col
}
