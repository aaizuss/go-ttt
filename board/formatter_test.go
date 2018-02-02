package board

import (
	"reflect"
	"testing"
)

func TestFormattedString(t *testing.T) {
	board := New(3)
	expected := " 0 | 1 | 2\n-----------\n 3 | 4 | 5\n-----------\n 6 | 7 | 8\n"

	result := board.FormattedString()

	if !reflect.DeepEqual(expected, result) {
		t.Errorf("Expected\n%v, got\n%v", expected, result)
	}
}

func TestConvertRowsToStrings(t *testing.T) {
	board := New(3)
	expected := []string{" 0 | 1 | 2", " 3 | 4 | 5", " 6 | 7 | 8"}

	result := board.convertRowsToStrings()

	if !reflect.DeepEqual(expected, result) {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

func TestRowToString(t *testing.T) {
	row := []string{"0", "1", "x"}
	expected := " 0 | 1 | x"

	result := rowToString(row)

	if result != expected {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

func TestIndexedRowsFromNewBoard(t *testing.T) {
	board := New(3)
	want := [][]string{
		{"0", "1", "2"},
		{"3", "4", "5"},
		{"6", "7", "8"},
	}

	result := board.indexedRows()

	if !reflect.DeepEqual(want, result) {
		t.Errorf("Expected %v, got %v", want, result)
	}
}

func TestIndexedRowsFromMarkedBoard(t *testing.T) {
	board := New(3)
	board.MarkSpace(1, "x")
	board.MarkSpace(4, "x")
	board.MarkSpace(7, "x")

	want := [][]string{
		{"0", "x", "2"},
		{"3", "x", "5"},
		{"6", "x", "8"},
	}
	result := board.indexedRows()

	if !reflect.DeepEqual(want, result) {
		t.Errorf("Expected %v, got %v", want, result)
	}
}
