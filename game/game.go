package game

import (
	"github.com/aaizuss/tictactoe/ai"
	"github.com/aaizuss/tictactoe/board"
	"github.com/aaizuss/tictactoe/console"
)

type ReadWriter interface {
	Read() string
	Write(message string)
}

type UIReadWriter interface {
	ReadWriter
	UI
}

type UI interface {
	ShowBoard(board board.Board)
	ShowOutcome(board board.Board)
	GetMove(board board.Board) (move int)
	Show(key string)
}

type Game struct {
	board   board.Board
	players []Player
	ui      UIReadWriter
}

func New() *Game {
	players := aiHuman()

	return &Game{
		board:   board.New(3),
		players: players,
		ui:      console.New(),
	}
}

func (game *Game) Play() {
	board := game.board
	ui := game.ui
	players := game.players

	ui.Show("welcome")
	ui.ShowBoard(board)

	for !board.GameOver() {
		//move := ui.GetMove(board)
		move := game.GetMove()
		board.MarkSpace(move, players[0].marker)
		game.TogglePlayer()
		ui.ShowBoard(board)
	}

	ui.ShowOutcome(board)
}

func (game *Game) GetMove() int {
	var move int
	if game.currentPlayer().isHuman {
		move = game.ui.GetMove(game.board)
	} else {
		move = ai.ChooseMove(game.board, game.markers(), game.currentPlayer().marker)
	}
	return move
}

func (game *Game) TogglePlayer() {
	game.players[0], game.players[1] = game.players[1], game.players[0]
}

func (game *Game) currentPlayer() Player {
	return game.players[0]
}

func (game *Game) markers() []string {
	return []string{game.players[0].marker, game.players[1].marker}
}
