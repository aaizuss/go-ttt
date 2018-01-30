package board_test

import (
	"github.com/aaizuss/tictactoe/board"
	"reflect"

	"testing"
)

func TestWinningCombos(t *testing.T) {
	board := boardWithNumberMarks()

	expectedWinningCombos := [][]string{
		{"0", "1", "2"},
		{"3", "4", "5"},
		{"6", "7", "8"},
		{"0", "3", "6"},
		{"1", "4", "7"},
		{"2", "5", "8"},
		{"0", "4", "8"},
		{"2", "4", "6"},
	}

	result := board.WinningCombos()

	if !reflect.DeepEqual(expectedWinningCombos, result) {
		t.Errorf("Expected %v, got %v", expectedWinningCombos, result)
	}
}

func TestIsFullReturnsTrueWhenBoardIsFull(t *testing.T) {
	board := boardWithNumberMarks()
	expected := true

	result := board.IsFull()

	if result != expected {
		t.Errorf("Expected true, got %v", result)
	}
}

func TestIsFullReturnsFalseWhenBoardIsNotFull(t *testing.T) {
	notFullBoards := []board.Board{nearlyFullBoard(), board.New(3)}

	for _, board := range notFullBoards {
		result := board.IsFull()
		if result != false {
			t.Errorf("Expected false, got %v", result)
		}
	}
}

func TestHasWinnerDiagonal(t *testing.T) {
	board := board.New(3)
	board.MarkSpace(0, "x")
	board.MarkSpace(4, "x")
	board.MarkSpace(8, "x")

	expected := true
	result := board.HasWinner()

	if result != expected {
		t.Errorf("Expected true, got %v", result)
	}
}

func TestHasWinnerRow(t *testing.T) {
	board := board.New(3)
	board.MarkSpace(3, "x")
	board.MarkSpace(4, "x")
	board.MarkSpace(5, "x")

	expected := true
	result := board.HasWinner()

	if result != expected {
		t.Errorf("Expected true, got %v", result)
	}
}

func TestHasWinnerColumn(t *testing.T) {
	board := board.New(3)
	board.MarkSpace(2, "x")
	board.MarkSpace(5, "x")
	board.MarkSpace(8, "x")

	expected := true
	result := board.HasWinner()

	if result != expected {
		t.Errorf("Expected true, got %v", result)
	}
}

func TestHasWinnerReturnsFalseWhenNoWinner(t *testing.T) {
	board := board.New(3)
	board.MarkSpace(2, "x")
	board.MarkSpace(5, "x")
	board.MarkSpace(8, "o")

	expected := false
	result := board.HasWinner()

	if result != expected {
		t.Errorf("Expected false, got %v", result)
	}
}

func TestWinningMarker(t *testing.T) {
	board := board.New(3)
	board.MarkSpace(0, "x")
	board.MarkSpace(4, "x")
	board.MarkSpace(8, "x")

	expectedWinner := "x"
	result, _ := board.Winner()

	if result != expectedWinner {
		t.Errorf("Expected x, got %v", result)
	}
}

func TestWinningMarkerReturnsNothingWhenNoWinner(t *testing.T) {
	board := board.New(3)
	board.MarkSpace(0, "x")
	board.MarkSpace(4, "x")
	board.MarkSpace(8, "o")

	result, err := board.Winner()

	if err == nil {
		t.Errorf("Expected NoWinnerError, got %v", result)
	}
}

func TestIsTieReturnsTrueWhenThereIsATie(t *testing.T) {
	board := tieBoard()

	expected := true
	result := board.IsTie()

	if result != expected {
		t.Errorf("Expected true, got %v", result)
	}
}

func TestIsTieReturnsFalseWhenNoTie(t *testing.T) {
	board := nearlyFullBoard()

	expected := false
	result := board.IsTie()

	if result != expected {
		t.Errorf("Expected false, got %v", result)
	}
}

func nearlyFullBoard() board.Board {
	board := board.New(3)

	for i := 0; i < 8; i++ {
		board.MarkSpace(i, "x")
	}
	return board
}

func tieBoard() board.Board {
	board := board.New(3)

	board.MarkSpace(0, "x")
	board.MarkSpace(1, "x")
	board.MarkSpace(2, "o")
	board.MarkSpace(3, "o")
	board.MarkSpace(4, "o")
	board.MarkSpace(5, "x")
	board.MarkSpace(6, "x")
	board.MarkSpace(7, "o")
	board.MarkSpace(8, "x")
	return board
}
