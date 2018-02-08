package console

import (
	"testing"

	"github.com/aaizuss/tictactoe/board"
)

func TestIsValidMoveChoice(t *testing.T) {
	board := board.New(3)
	board.MarkSpace(2, "x")

	tests := []struct {
		moveChoice       string
		expectedValidity bool
	}{
		{"2", false},
		{"1", true},
		{"8", true},
		{"9", false},
		{"a", false},
		{"abc", false},
	}

	for _, test := range tests {
		validity := IsValidMoveChoice(board, test.moveChoice)
		if validity != test.expectedValidity {
			t.Errorf("For board %v, expected IsValidMoveChoice(%s) to be %v, got %v", board, test.moveChoice, test.expectedValidity, validity)
		}
	}
}

func TestIsValidGameChoiceReturnsTrue(t *testing.T) {
	input := "2"

	expect := true
	result := IsValidGameChoice(input)

	if result != expect {
		t.Errorf("Expected %v, got %v", expect, result)
	}
}
