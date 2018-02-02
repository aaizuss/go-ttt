// +build aitest

package game

import (
	"github.com/aaizuss/tictactoe/board"
	"testing"
)

func TestComputerVsComputerEndsInTie(t *testing.T) {
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
