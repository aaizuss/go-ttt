package game

import (
	"reflect"
	"strings"

	"github.com/aaizuss/tictactoe/board"

	"testing"
)

func emptyPlayers() []Player {
	return []Player{Player{}, Player{}}
}

func TestSetupPlayers(t *testing.T) {
	cli := MockConsole{UserInput: "1"}
	game := Game{board: board.New(3), players: emptyPlayers(), ui: &cli}

	expected := []Player{
		Player{marker: "x", isHuman: true},
		Player{marker: "o", isHuman: true},
	}

	game.SetupPlayers()
	result := game.players

	if !reflect.DeepEqual(expected, result) {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

func TestTogglePlayer(t *testing.T) {
	cli := MockConsole{}
	players := humanHuman()
	game := Game{board: board.New(3), players: players, ui: &cli}

	expected := []Player{
		Player{marker: "o", isHuman: true},
		Player{marker: "x", isHuman: true},
	}

	game.togglePlayer()
	result := game.players

	if !reflect.DeepEqual(expected, result) {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

func TestTakeTurnsWhenAboutToTie(t *testing.T) {
	cli := MockConsole{UserInput: "8"}
	players := humanHuman()
	game := Game{board: almostTieBoard(), players: players, ui: &cli}

	game.takeTurns()
	expected := "It's a tie!\n"
	result := cli.Output

	if !strings.Contains(result, expected) {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

func TestTakeTurnsHaltsWhenThereIsAWinner(t *testing.T) {
	cli := MockConsole{UserInput: "2"}
	players := humanHuman()
	game := Game{board: xAlmostWinBoard(), players: players, ui: &cli}

	game.takeTurns()
	expected := "win"
	result := cli.Output

	if !strings.Contains(result, expected) {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

func TestComputersAgainstEachOtherResultsInTie(t *testing.T) {
	cli := MockConsole{}
	p1 := Player{marker: "x", isHuman: false}
	p2 := Player{marker: "o", isHuman: false}
	players := []Player{p1, p2}
	game := Game{board: board.New(3), players: players, ui: &cli}

	game.takeTurns()
	result := game.board.IsTie()
	expect := true

	if result != expect {
		t.Errorf("Expected %v, got %v", expect, result)
	}
}

func almostTieBoard() board.Board {
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
