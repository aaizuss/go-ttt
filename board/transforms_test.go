package board_test

import (
	"github.com/aaizuss/tictactoe/board"
	"reflect"
	"strconv"

	"testing"
)

func boardWithNumberMarks() board.Board {
	board := board.New(3)

	for index := range board.Spaces() {
		board.MarkSpace(index, strconv.Itoa(index))
	}

	return board
}

func TestRowsGetsRows(t *testing.T) {
	board := boardWithNumberMarks()

	rows := [][]string{
		{"0", "1", "2"},
		{"3", "4", "5"},
		{"6", "7", "8"},
	}

	if !reflect.DeepEqual(rows, board.Rows()) {
		t.Errorf("Expected %v, got %v", rows, board.Rows())
	}
}

func TestColsGetsColumns(t *testing.T) {
	board := boardWithNumberMarks()

	cols := [][]string{
		{"0", "3", "6"},
		{"1", "4", "7"},
		{"2", "5", "8"},
	}

	if !reflect.DeepEqual(cols, board.Cols()) {
		t.Errorf("Expected %v, got %v", cols, board.Cols())
	}
}

func TestDiagonalsGetsDiagonals(t *testing.T) {
	board := boardWithNumberMarks()

	diagonals := [][]string{
		{"0", "4", "8"},
		{"2", "4", "6"},
	}

	if !reflect.DeepEqual(diagonals, board.Diagonals()) {
		t.Errorf("Expected %v, got %v", diagonals, board.Diagonals())
	}
}
