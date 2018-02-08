package game

import (
	"reflect"

	"github.com/aaizuss/go-ttt/board"

	"testing"
)

func emptyPlayers() []Player {
	return []Player{Player{}, Player{}}
}

func TestSetupPlayers(t *testing.T) {
	cli := MockConsole{UserInput: "1"}
	game := Game{board: board.New(3), players: emptyPlayers(), ui: &cli}

	expect := []Player{
		Player{marker: "x", isHuman: true},
		Player{marker: "o", isHuman: true},
	}

	game.SetupPlayers()
	result := game.players

	if !reflect.DeepEqual(expect, result) {
		t.Errorf("Expected %v, got %v", expect, result)
	}
}

func TestTogglePlayer(t *testing.T) {
	cli := MockConsole{}
	players := humanHuman()
	game := Game{board: board.New(3), players: players, ui: &cli}

	expect := []Player{
		Player{marker: "o", isHuman: true},
		Player{marker: "x", isHuman: true},
	}

	game.togglePlayer()
	result := game.players

	if !reflect.DeepEqual(expect, result) {
		t.Errorf("Expected %v, got %v", expect, result)
	}
}

func TestTakeTurnsOnTie(t *testing.T) {
	cli := MockConsole{UserInput: "8"}
	players := humanHuman()
	game := Game{board: almostTieBoard(), players: players, ui: &cli}

	game.takeTurns()
	expect := true
	result := game.board.IsTie()

	if result != expect {
		t.Errorf("Expected IsTie to be %v, got %v", expect, result)
	}
}

func TestTakeTurnsHaltsWhenThereIsAWinner(t *testing.T) {
	cli := MockConsole{UserInput: "2"}
	players := humanHuman()
	game := Game{board: xAlmostWinBoard(), players: players, ui: &cli}

	game.takeTurns()
	expectedWinner := "x"
	result, _ := game.board.Winner()

	if result != expectedWinner {
		t.Errorf("Expected winner %v, got %v", expectedWinner, result)
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
