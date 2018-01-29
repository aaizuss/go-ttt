package board

// import "fmt"

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
	cols := make([][]string, width)

	for i := 0; i < width; i++ {
		var col []string

		for j := 0; j < width; j++ {
			col = append(col, board.spaces[j*width+i])
		}
		cols[i] = col
	}
	return cols
}

func (board Board) Diagonals() [][]string {
	rows := board.Rows()
	diagonals := make([][]string, 2)

	diagonals[0] = append(diagonals[0], rows[0][0], rows[1][1], rows[2][2])
	diagonals[1] = append(diagonals[1], rows[0][2], rows[1][1], rows[2][0])

	return diagonals
}
