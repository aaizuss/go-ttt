package board_test

import (
	"reflect"

	"github.com/aaizuss/tictactoe/board"

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
	expected := false

	for _, board := range notFullBoards {
		isFull := board.IsFull()
		if isFull != expected {
			t.Errorf("Expected %v, got %v", expected, isFull)
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

func TestWinner(t *testing.T) {
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

func TestWinnerReturnsErrorWhenNoWinner(t *testing.T) {
	aBoard := board.New(3)
	aBoard.MarkSpace(0, "x")
	aBoard.MarkSpace(4, "x")
	aBoard.MarkSpace(8, "o")

	winner, err := aBoard.Winner()

	if err != board.NoWinnerError {
		t.Errorf("Expected NoWinnerError, got %v", winner)
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

func TestGameOverReturnsFalseForOngoingGame(t *testing.T) {
	board := ongoingGame()

	expected := false
	result := board.GameOver()

	if result != expected {
		t.Errorf("Expected %v, got %v\nfrom board: %v", expected, result, board)
	}
}

func TestIsEmptyReturnsFalseForOngoingGame(t *testing.T) {
	board := ongoingGame()

	expected := false
	result := board.IsEmpty()

	if result != expected {
		t.Errorf("Expected %v, got %v\nfrom board: %v", expected, result, board)
	}
}

func TestIsEmptyReturnsTrueForNewBoard(t *testing.T) {
	board := board.New(3)

	expected := true
	result := board.IsEmpty()

	if result != expected {
		t.Errorf("Expected %v, got %v\nfrom board: %v", expected, result, board)
	}
}

func nearlyFullBoard() board.Board {
	board := board.New(3)

	for i := 0; i < 8; i++ {
		board.MarkSpace(i, "x")
	}
	return board
}

func ongoingGame() board.Board {
	board := board.New(3)

	board.MarkSpace(0, "x")
	board.MarkSpace(1, "x")
	board.MarkSpace(2, "o")
	board.MarkSpace(3, "o")
	board.MarkSpace(4, "o")
	board.MarkSpace(5, "x")
	board.MarkSpace(6, "x")
	board.MarkSpace(7, "o")
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
