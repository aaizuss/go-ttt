package game

import (
	"reflect"
	"strings"

	"github.com/aaizuss/tictactoe/board"

	"testing"
)

func TestTogglePlayer(t *testing.T) {
	cli := MockConsole{}
	game := Game{board: board.New(3), players: []string{"x", "o"}, ui: &cli}

	expected := []string{"o", "x"}
	game.TogglePlayer()
	result := game.players

	if !reflect.DeepEqual(expected, result) {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

func TestPlayWhenAboutToTie(t *testing.T) {
	cli := MockConsole{UserInput: "8"}
	game := Game{board: almostTieBoard(), players: []string{"x", "o"}, ui: &cli}

	game.Play()
	expected := "It's a tie!\n"
	result := cli.Output

	if !strings.Contains(result, expected) {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

func TestPlayHaltsWhenThereIsAWinner(t *testing.T) {
	cli := MockConsole{UserInput: "2"}
	game := Game{board: xAlmostWinBoard(), players: []string{"x", "o"}, ui: &cli}

	game.Play()
	expected := "win"
	result := cli.Output

	if !strings.Contains(result, expected) {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

func almostTieBoard() board.Board {
	// x will move to 8 to tie
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

func xAlmostWinBoard() board.Board {
	board := board.New(3)
	board.MarkSpace(0, "x")
	board.MarkSpace(1, "x")
	return board
}
