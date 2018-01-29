package board_test

import (
	"github.com/aaizuss/tictactoe"
	"testing"
)

func TestNewBoardReturnsBoardFromDimensions(t *testing.T) {
	blankBoard := board.New(3)

	if blankBoard.NumSpaces() != 9 {
		t.Errorf("Expected length 9, got %d", blankBoard.NumSpaces)
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

func TestMarkSpace(t *testing.T) {
	myboard := board.New(3)

	myboard.MarkSpace(6, "x")

	if myboard.GetSpace(6) != "x" {
		t.Errorf("Expected space 6 to be x, got %q", myboard.GetSpace(6))
	}
}
