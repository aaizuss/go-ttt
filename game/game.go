package game

import (
	"github.com/aaizuss/tictactoe/board"
	"github.com/aaizuss/tictactoe/console"
)

type Game struct {
	board   board.Board
	players []string
	ui      console.UIReadWriter //lol this reads poorly
}

func New() Game {
	return Game{board: board.New(3), players: []string{"x", "o"}, ui: console.New()}
}

func (game Game) IsOver() bool {
	return game.board.IsTie() || game.board.HasWinner()
}

func (game Game) TogglePlayer() {
	current, next := game.players[1], game.players[0]
	game.players[0] = current
	game.players[1] = next
}

func (game Game) Play() {
	board := game.board
	ui := game.ui
	players := game.players

	ui.Show("welcome")
	ui.ShowBoard(board)

	for !game.IsOver() {
		move := ui.GetMove(board)
		board.MarkSpace(move, players[0])
		game.TogglePlayer()
		ui.ShowBoard(board)
	}

	ui.ShowOutcome(board)
}
