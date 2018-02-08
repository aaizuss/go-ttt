package board_test

import (
	"reflect"
	"testing"

	"github.com/aaizuss/tictactoe/board"
)

func TestNewBoardCalculatesNumSpacesFromDimensions(t *testing.T) {
	blankBoard := board.New(3)

	expect := 9

	if blankBoard.NumSpaces != expect {
		t.Errorf("Expected %d, got %d", expect, blankBoard.NumSpaces)
	}
}

func TestNewBoardReturnsEmptyBoard(t *testing.T) {
	blankBoard := board.New(3)
	expectedSpaces := []string{
		"_", "_", "_",
		"_", "_", "_",
		"_", "_", "_",
	}
	gotSpaces := blankBoard.Spaces()

	if !reflect.DeepEqual(expectedSpaces, gotSpaces) {
		t.Errorf("Expected board to be %v, got %v", expectedSpaces, gotSpaces)
	}
}

func TestMarkSpace(t *testing.T) {
	myboard := board.New(3)
	expect := "x"

	myboard.MarkSpace(6, "x")

	if myboard.Spaces()[6] != expect {
		t.Errorf("Expected space 6 to be %v, got %v", expect, myboard.Spaces()[6])
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

	expect := []int{0, 1, 3, 5, 6, 7, 8}
	result := board.EmptySpaces()

	if !reflect.DeepEqual(expect, result) {
		t.Errorf("Expected %v, got %v", expect, result)
	}
}

func TestResetSpace(t *testing.T) {
	testBoard := board.New(3)
	testBoard.MarkSpace(4, "x")

	testBoard.ResetSpace(4)
	spaceAfterReset := testBoard.Spaces()[4]

	if spaceAfterReset != board.EmptySpace {
		t.Errorf("Expected space 4 to be empty, got %v", spaceAfterReset)
	}
}
