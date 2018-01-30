package console

import (
	"github.com/aaizuss/tictactoe/board"
	"strconv"
)

func (cli CommandLine) GetMove(board board.Board) (move int) {
	for {
		cli.Show("choose-space")
		move := cli.Read()

		if IsValidMoveChoice(board, move) {
			return toInt(move)
		}
	}
}

func toInt(inputMove string) int {
	move, _ := strconv.Atoi(inputMove)
	return move
}
