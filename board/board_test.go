package board_test

import (
	"github.com/aaizuss/tictactoe/board"
	"testing"
)

func TestNewBoardReturnsBoardFromDimensions(t *testing.T) {
	blankBoard := board.New(3)

	expected := 9

	if blankBoard.NumSpaces != expected {
		t.Errorf("Expected num spaces to %d, got %d", expected, blankBoard.NumSpaces)
	}
}

func TestNewBoardReturnsEmptyBoard(t *testing.T) {
	blankBoard := board.New(3)
	expected := "_"

	for i, space := range blankBoard.Spaces() {
		if space != expected {
			t.Errorf("Expected board[%d] to be '_', got %q", i, space)
		}
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

func TestSpaceExistsIsTrueWhenSpaceExists(t *testing.T) {
	board := board.New(3)
	expected := true

	validSpaces := []int{
		0, 1, 2, 3, 4, 5, 6, 7, 8,
	}

	for _, validSpaceIndex := range validSpaces {
		result := board.SpaceExists(validSpaceIndex)
		if result != expected {
			t.Errorf("Expected SpaceExists(%d) to be true, got %v", validSpaceIndex, result)
		}
	}
}

func TestSpaceExistsIsFalseWhenSpaceDoesNotExist(t *testing.T) {
	board := board.New(3)
	expected := false

	invalidSpaces := []int{
		-1, 9, -8, 10,
	}

	for _, invalidSpaceIndex := range invalidSpaces {
		result := board.SpaceExists(invalidSpaceIndex)
		if result != expected {
			t.Errorf("Expected SpaceExists(%d) to be false, got %v", invalidSpaceIndex, result)
		}
	}
}

func TestIsValidMoveReturnsTrueWhenSpaceIsAvailable(t *testing.T) {
	board := board.New(3)
	board.MarkSpace(0, "x")
	choice := 5

	expected := true
	result := board.IsValidMove(choice)

	if result != expected {
		t.Errorf("Expected %d to be a valid move, got %v", choice, result)
	}
}

func TestIsValidMoveReturnsFalseWhenTheSpaceDoesNotExist(t *testing.T) {
	board := board.New(3)
	choice := 9

	expected := false
	result := board.IsValidMove(choice)

	if result != expected {
		t.Errorf("Expected %d to be %v, got %v", choice, expected, result)
	}
}

func TestIsValidMoveReturnsFalseWhenTheSpaceIsMarked(t *testing.T) {
	board := board.New(3)
	board.MarkSpace(0, "x")
	choice := 0

	expected := false
	result := board.IsValidMove(choice)

	if result != expected {
		t.Errorf("Expected %d to be %v, got %v", choice, expected, result)
	}
}
