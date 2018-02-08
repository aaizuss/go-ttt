package game

import (
	"strconv"

	"github.com/aaizuss/go-ttt/board"
	"github.com/aaizuss/go-ttt/console"
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
		move := cli.Read()
		if console.IsValidMoveChoice(board, move) {
			move, _ := strconv.Atoi(move)
			return move
		} else {
			cli.Show("invalid move")
		}
	}
}

func (cli *MockConsole) GetGameChoice() string {
	for {
		cli.Show("game menu")
		choice := cli.Read()

		if console.IsValidGameChoice(choice) {
			return choice
		} else {
			cli.Show("invalid choice")
		}
	}
}

func (cli *MockConsole) Show(message string) {
	cli.Write(message)
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
	cli.Show("move recap")
}
