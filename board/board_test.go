package board_test

import (
	"reflect"
	"testing"

	"github.com/aaizuss/tictactoe/board"
)

func TestNewBoardReturnsBoardFromDimensions(t *testing.T) {
	blankBoard := board.New(3)

	expected := 9

	if blankBoard.NumSpaces != expected {
		t.Errorf("Expected num spaces to %d, got %d", expected, blankBoard.NumSpaces)
	}
}

func TestNewBoardReturnsEmptyBoard(t *testing.T) {
	board := board.New(3)
	expectedSpaces := []string{
		"_", "_", "_",
		"_", "_", "_",
		"_", "_", "_",
	}
	spaces := board.Spaces()

	if !reflect.DeepEqual(expectedSpaces, spaces) {
		t.Errorf("Expected board to be %v, got %v", expectedSpaces, spaces)
	}
}

func TestMarkSpace(t *testing.T) {
	myboard := board.New(3)
	expected := "x"

	myboard.MarkSpace(6, "x")

	if myboard.Spaces()[6] != expected {
		t.Errorf("Expected space 6 to be %q, got %q", expected, myboard.Spaces()[6])
	}
}

func TestIsValidMoveChecksSpaceExistsAndNotMarked(t *testing.T) {
	board := board.New(3)
	board.MarkSpace(0, "x")
	board.MarkSpace(4, "o")

	tests := []struct {
		space            int
		expectedValidity bool
	}{
		{0, false},
		{1, true},
		{2, true},
		{3, true},
		{4, false},
		{9, false},
		{-1, false},
	}
	for _, test := range tests {
		result := board.IsValidMove(test.space)
		if result != test.expectedValidity {
			t.Errorf("expected IsValidMove(%d) to be %v, got %v", test.space, test.expectedValidity, result)
		}
	}
}

func TestEmptySpaces(t *testing.T) {
	board := board.New(3)
	board.MarkSpace(4, "x")
	board.MarkSpace(2, "o")

	expected := []int{0, 1, 3, 5, 6, 7, 8}
	result := board.EmptySpaces()

	if !reflect.DeepEqual(expected, result) {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}
