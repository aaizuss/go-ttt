package board_test

import (
	"reflect"
	"strconv"

	"github.com/aaizuss/go-ttt/board"

	"testing"
)

func TestRowsGetsRows(t *testing.T) {
	board := boardWithNumberMarks()
	expectedRows := [][]string{
		{"0", "1", "2"},
		{"3", "4", "5"},
		{"6", "7", "8"},
	}

	result := board.Rows()

	if !reflect.DeepEqual(expectedRows, result) {
		t.Errorf("Expected %v, got %v", expectedRows, result)
	}
}

func TestColsGetsColumns(t *testing.T) {
	board := boardWithNumberMarks()
	expectedCols := [][]string{
		{"0", "3", "6"},
		{"1", "4", "7"},
		{"2", "5", "8"},
	}

	result := board.Cols()

	if !reflect.DeepEqual(expectedCols, result) {
		t.Errorf("Expected %v, got %v", expectedCols, result)
	}
}

func TestDiagonalsGetsDiagonals(t *testing.T) {
	board := boardWithNumberMarks()
	expectedDiagonals := [][]string{
		{"0", "4", "8"},
		{"2", "4", "6"},
	}

	result := board.Diagonals()

	if !reflect.DeepEqual(expectedDiagonals, result) {
		t.Errorf("Expected %v, got %v", expectedDiagonals, result)
	}
}

func boardWithNumberMarks() board.Board {
	board := board.New(3)

	for index := range board.Spaces() {
		board.MarkSpace(index, strconv.Itoa(index))
	}

	return board
}
