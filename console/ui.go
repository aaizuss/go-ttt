package console

import (
	"github.com/aaizuss/tictactoe/board"
	"strconv"
)

func (cli CommandLine) DisplayBoard(board board.Board) {
	cli.Write(board.FormattedString())
}

func (cli CommandLine) GetMove(board board.Board) (move int) {
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
