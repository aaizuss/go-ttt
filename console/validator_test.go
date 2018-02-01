package console

import (
	"testing"

	"github.com/aaizuss/tictactoe/board"
)

func TestIsValidMoveChoiceReturnsTrue(t *testing.T) {
	board := board.New(3)
	board.MarkSpace(0, "x")
	inputMove := "1"

	expected := true
	result := IsValidMoveChoice(board, inputMove)

	if result != expected {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

func TestIsValidMoveChoiceReturnsFalseWhenChoiceIsInvalid(t *testing.T) {
	board := board.New(3)
	board.MarkSpace(2, "x")

	invalidChoices := []string{
		"-1", "9", "2", "a", "abc",
	}
	expected := false

	for _, invalidChoice := range invalidChoices {
		result := IsValidMoveChoice(board, invalidChoice)
		if result != expected {
			t.Errorf("Expected IsValidMoveChoice(%v) to be %v, got %v", invalidChoice, expected, result)
		}
	}
}

func TestIsValidGameChoiceReturnsTrue(t *testing.T) {
	input := "2"

	expected := true
	result := IsValidGameChoice(input)

	if result != expected {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}
