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
	ShowMoveRecap(marker string, move int)
	GetMove(board board.Board) (move int)
	GetGameChoice() string
	Show(key string)
}

type Game struct {
	board   board.Board
	players []Player
	ui      UIReadWriter
}

func New() *Game {
	blankPlayers := []Player{Player{}, Player{}}
	return &Game{
		board:   board.New(3),
		players: blankPlayers,
		ui:      console.New(),
	}
}

func (game *Game) Play() {
	game.ui.Show("welcome")
	game.SetupPlayers()
	game.ui.ShowBoard(game.board)
	game.takeTurns()
}

func (game *Game) takeTurns() {
	board := game.board
	ui := game.ui

	for !board.GameOver() {
		currentPlayer := game.players[0]
		move := game.getMove()
		board.MarkSpace(move, currentPlayer.marker)
		ui.ShowBoard(board)
		ui.ShowMoveRecap(currentPlayer.marker, move)
		game.togglePlayer()
	}

	ui.ShowOutcome(board)
}

func (game *Game) getMove() int {
	var move int
	if game.currentPlayer().isHuman {
		move = game.ui.GetMove(game.board)
	} else {
		move = ai.ChooseMove(game.board, game.markers(), game.currentPlayer().marker)
	}
	return move
}

func (game *Game) togglePlayer() {
	game.players[0], game.players[1] = game.players[1], game.players[0]
}

func (game *Game) currentPlayer() Player {
	return game.players[0]
}

func (game *Game) markers() []string {
	return []string{game.players[0].marker, game.players[1].marker}
}
