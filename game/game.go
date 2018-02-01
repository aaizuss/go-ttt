package game

import (
	"github.com/aaizuss/tictactoe/ai"
	"github.com/aaizuss/tictactoe/board"
	"github.com/aaizuss/tictactoe/console"
	"strconv"
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
	game.SetupPlayers()
	game.takeTurns()
}

func (game *Game) takeTurns() {
	board := game.board
	ui := game.ui
	players := game.players

	ui.ShowBoard(board)

	for !board.GameOver() {
		move := game.GetMove()
		board.MarkSpace(move, players[0].marker)
		ui.ShowBoard(board)
		ui.Write(afterMoveMessage(players[0], move))
		game.TogglePlayer()
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

// todo: move these and figure out how to incorporate into Messages hash
func afterMoveMessage(player Player, move int) string {
	choice := strconv.Itoa(move)
	return player.marker + " marked " + choice + "\n"
}

func computerMessage(player Player, move int) string {
	choice := strconv.Itoa(move)
	return "\n" + player.marker + " marked " + choice + "\n\n"
}
