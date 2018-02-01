package console

import (
	"github.com/aaizuss/tictactoe/board"
	"strconv"
)

func (cli *Console) ShowBoard(board board.Board) {
	cli.Write(wrapNewLine(board.FormattedString()))
}

func (cli *Console) Show(key string) {
	cli.Write(Messages[key])
}

func (cli *Console) ShowOutcome(board board.Board) {
	if board.IsTie() {
		cli.Show("tie")
	} else {
		winner, _ := board.Winner()
		winMessage := "Player " + winner + " wins!\n"
		cli.Write(winMessage)
	}
}

func (cli *Console) GetMove(board board.Board) (move int) {
	for {
		cli.Show("choose-space")
		move := cli.Read()

		if IsValidMoveChoice(board, move) {
			return toInt(move)
		} else {
			cli.Show("invalid-move")
		}
	}
}

func (cli *Console) GetGameChoice() string {
	for {
		cli.Show("choose-game")
		choice := cli.Read()

		if IsValidGameChoice(choice) {
			return choice
		} else {
			cli.Show("invalid-choice")
		}
	}
}

func toInt(inputMove string) int {
	move, _ := strconv.Atoi(inputMove)
	return move
}

// put in a json file at some point?
var Messages = map[string]string{
	"welcome":        "|----------------------------|\n|-- Welcome to Tic Tac Toe --|\n|----------------------------|\n\n",
	"tie":            "It's a tie!\n",
	"choose-game":    "Choose a game type.\n1  human v human\n2  human v computer\n3  computer v human\n\n",
	"choose-space":   "Enter a number 0-8 to mark that position on the board: ",
	"invalid-move":   "You can't move there. ",
	"invalid-choice": "That's not an option.\n",
	"taken-space":    "That space is taken. ",
}

func wrapNewLine(message string) string {
	return "\n" + message + "\n"
}
