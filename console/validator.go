package console

import (
	"github.com/aaizuss/tictactoe/board"
	"strconv"
)

func IsValidMoveChoice(board board.Board, inputMove string) bool {
	move, err := strconv.Atoi(inputMove)
	if err != nil {
		return false
	}
	return board.IsValidMove(move)
}
