package game

import (
	"github.com/aaizuss/go-ttt/ai"
	"github.com/aaizuss/go-ttt/board"
	"github.com/aaizuss/go-ttt/console"
)

type UI interface {
	ShowBoard(board board.Board)
	ShowOutcome(board board.Board)
	ShowMoveRecap(marker string, move int)
	GetMove(board board.Board) (move int)
	GetGameChoice() string
	Show(message string)
}

type Game struct {
	board   board.Board
	players []Player
	ui      UI
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
	game.SetupPlayers()
	game.ShowInitialBoard()
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

func (game *Game) ShowInitialBoard() {
	if game.currentPlayer().isHuman {
		game.ui.ShowBoard(game.board)
	}
}
