package board_test

import (
	"reflect"

	"github.com/aaizuss/go-ttt/board"

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
	expect := true

	result := board.IsFull()

	if result != expect {
		t.Errorf("Expected %v, got %v", expect, result)
	}
}

func TestIsFullReturnsFalseWhenBoardIsNotFull(t *testing.T) {
	notFullBoards := []board.Board{nearlyFullBoard(), board.New(3)}
	expect := false

	for _, board := range notFullBoards {
		isFull := board.IsFull()
		if isFull != expect {
			t.Errorf("Expected %v, got %v", expect, isFull)
		}
	}
}

func TestHasWinnerDiagonal(t *testing.T) {
	board := board.New(3)
	board.MarkSpace(0, "x")
	board.MarkSpace(4, "x")
	board.MarkSpace(8, "x")

	expect := true
	result := board.HasWinner()

	if result != expect {
		t.Errorf("Expected %v, got %v", expect, result)
	}
}

func TestHasWinnerRow(t *testing.T) {
	board := board.New(3)
	board.MarkSpace(3, "x")
	board.MarkSpace(4, "x")
	board.MarkSpace(5, "x")

	expect := true
	result := board.HasWinner()

	if result != expect {
		t.Errorf("Expected %v, got %v", expect, result)
	}
}

func TestHasWinnerColumn(t *testing.T) {
	board := board.New(3)
	board.MarkSpace(2, "x")
	board.MarkSpace(5, "x")
	board.MarkSpace(8, "x")

	expect := true
	result := board.HasWinner()

	if result != expect {
		t.Errorf("Expected %v, got %v", expect, result)
	}
}

func TestHasWinnerReturnsFalseWhenNoWinner(t *testing.T) {
	board := board.New(3)
	board.MarkSpace(2, "x")
	board.MarkSpace(5, "x")
	board.MarkSpace(8, "o")

	expect := false
	result := board.HasWinner()

	if result != expect {
		t.Errorf("Expected %v, got %v", expect, result)
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
		t.Errorf("Expected %v, got %v", expectedWinner, result)
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

	expect := true
	result := board.IsTie()

	if result != expect {
		t.Errorf("Expected %v, got %v", expect, result)
	}
}

func TestIsTieReturnsFalseWhenNoTie(t *testing.T) {
	board := nearlyFullBoard()

	expect := false
	result := board.IsTie()

	if result != expect {
		t.Errorf("Expected %v, got %v", expect, result)
	}
}

func TestGameOverReturnsFalseForOngoingGame(t *testing.T) {
	board := ongoingGame()

	expect := false
	result := board.GameOver()

	if result != expect {
		t.Errorf("Expected %v, got %v\nfrom board: %v", expect, result, board)
	}
}

func TestIsEmptyReturnsFalseForOngoingGame(t *testing.T) {
	board := ongoingGame()

	expect := false
	result := board.IsEmpty()

	if result != expect {
		t.Errorf("Expected %v, got %v\nfrom board: %v", expect, result, board)
	}
}

func TestIsEmptyReturnsTrueForNewBoard(t *testing.T) {
	board := board.New(3)

	expect := true
	result := board.IsEmpty()

	if result != expect {
		t.Errorf("Expected %v, got %v\nfrom board: %v", expect, result, board)
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
