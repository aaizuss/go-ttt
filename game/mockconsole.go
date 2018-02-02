package game

import (
	"strconv"

	"github.com/aaizuss/tictactoe/board"
	"github.com/aaizuss/tictactoe/console"
)

type MockConsole struct {
	UserInput string
	Output    string
}

func (cli MockConsole) Read() string {
	return cli.UserInput
}

func (cli *MockConsole) Write(message string) {
	cli.Output += message
}

func (cli *MockConsole) ShowBoard(board board.Board) {
	cli.Write(board.FormattedString())
}

func (cli *MockConsole) GetMove(board board.Board) (move int) {
	for {
		cli.Show("choose-space")
		move := cli.Read()

		if console.IsValidMoveChoice(board, move) {
			move, _ := strconv.Atoi(move)
			return move
		} else {
			cli.Show("invalid-move")
		}
	}
}

func (cli *MockConsole) GetGameChoice() string {
	for {
		cli.Show("choose-game")
		choice := cli.Read()

		if console.IsValidGameChoice(choice) {
			return choice
		} else {
			cli.Show("invalid-choice")
		}
	}
}

func (cli *MockConsole) Show(key string) {
	cli.Write(console.Messages[key])
}

func (cli *MockConsole) ShowOutcome(board board.Board) {
	if board.IsTie() {
		cli.Show("tie")
	} else {
		winner, _ := board.Winner()
		winMessage := "Player " + winner + " wins!\n"
		cli.Write(winMessage)
	}
}

func (cli *MockConsole) ShowMoveRecap(marker string, move int) {
	cli.Write(marker + " marked " + (strconv.Itoa(move)) + "\n")
}
