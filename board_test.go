package board_test

import (
	"github.com/aaizuss/tictactoe"
	"testing"
)

func TestNewBoardReturnsBoardFromDimensions(t *testing.T) {
	board := board.New(3)

	if len(board) != 9 {
		t.Errorf("Wanted length 9, got %d", len(board))
	}
}

func TestNewBoardReturnsEmptyBoard(t *testing.T) {
	blankBoard := board.New(3)

	for i := 0; i < len(blankBoard); i++ {
		if blankBoard[i] != board.EmptySpace {
			t.Errorf("Expected board[%d] to be '_', got %q", i, blankBoard[i])
		}
	}
}

func TestMarkSpace(t *testing.T) {
	blankBoard := board.New(3)

	result := board.MarkSpace(blankBoard, 6, "x")

	if result[6] != "x" {
		t.Errorf("Expected space 6 to be x, got %q", result[6])
	}
}
