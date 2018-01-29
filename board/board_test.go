package board_test

import (
	"github.com/aaizuss/tictactoe/board"
	"testing"
)

func TestNewBoardReturnsBoardFromDimensions(t *testing.T) {
	blankBoard := board.New(3)

	if blankBoard.NumSpaces != 9 {
		t.Errorf("Expected num spaces to be 9, got %d", blankBoard.NumSpaces)
	}
}

func TestNewBoardReturnsEmptyBoard(t *testing.T) {
	blankBoard := board.New(3)

	for i := range blankBoard.Spaces() {
		if blankBoard.IsMarked(i) {
			t.Errorf("Expected board[%d] to be '_', got %q", i, blankBoard.GetSpace(i))
		}
	}
}

func TestIsMarkedReturnsTrueWhenSpaceIsMarked(t *testing.T) {
	board := board.New(3)

	board.MarkSpace(4, "x")

	if !board.IsMarked(4) {
		t.Errorf("Expected IsMarked(4) to be true, got false")
	}
}

func TestIsMarkedReturnsFalseWhenSpaceIsEmpty(t *testing.T) {
	board := board.New(3)

	board.MarkSpace(4, "x")

	if board.IsMarked(3) {
		t.Errorf("Expected IsMarked(3) to be false, got true")
	}
}

func TestIsEmptyReturnsTrueWhenSpaceIsEmpty(t *testing.T) {
	board := board.New(3)

	board.MarkSpace(4, "x")

	if !board.IsEmpty(3) {
		t.Errorf("Expected IsEmpty(3) to be true, got false")
	}
}

func TestIsEmptyReturnsFalseWhenSpaceIsMarked(t *testing.T) {
	board := board.New(3)

	board.MarkSpace(4, "x")

	if board.IsEmpty(4) {
		t.Errorf("Expected IsEmpty(4) to be false, got true")
	}
}

func TestMarkSpace(t *testing.T) {
	myboard := board.New(3)

	myboard.MarkSpace(6, "x")

	if myboard.GetSpace(6) != "x" {
		t.Errorf("Expected space 6 to be x, got %q", myboard.GetSpace(6))
	}
}

func TestSpaceExistsIsTrueWhenSpaceExists(t *testing.T) {
	board := board.New(3)

	result := board.SpaceExists(3)

	if result != true {
		t.Errorf("expected true, got %q", result)
	}
}

func TestSpaceExistsIsFalseWhenSpaceDoesNotExist(t *testing.T) {
	board := board.New(3)

	result := board.SpaceExists(10)

	if result != false {
		t.Errorf("expected false, got %q", result)
	}
}
