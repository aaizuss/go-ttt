package console

import (
	"github.com/aaizuss/tictactoe/board"
	"strconv"
)

type UI interface {
	DisplayBoard(board board.Board)
	GetMove(board board.Board) (move int)
	Show(key string)
}

func (cli *CommandLine) DisplayBoard(board board.Board) {
	cli.Write(board.FormattedString())
}

func (cli *CommandLine) Show(key string) {
	cli.Write(messages[key])
}

func (cli *CommandLine) GetMove(board board.Board) (move int) {
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

func toInt(inputMove string) int {
	move, _ := strconv.Atoi(inputMove)
	return move
}

// put in a json file at some point?
var messages = map[string]string{
	"welcome":      "|----------------------------|\n|-- Welcome to Tic Tac Toe --|\n|----------------------------|\n",
	"tie":          "It's a tie!\n",
	"choose-space": "Enter a number 0-8 to mark that position on the board: ",
	"invalid-move": "You can't move there. ",
	"taken-space":  "That space is taken. ",
}
