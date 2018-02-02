package console

import (
	"strconv"

	"github.com/aaizuss/tictactoe/board"
)

func IsValidMoveChoice(board board.Board, inputMove string) bool {
	move, err := strconv.Atoi(inputMove)
	if err != nil {
		return false
	}
	return board.IsValidMove(move)
}

func IsValidGameChoice(input string) bool {
	return input == "1" || input == "2" || input == "3"
}
