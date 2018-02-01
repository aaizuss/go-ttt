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

func IsValidGameChoice(input string) bool {
	return input == "1" || input == "2" || input == "3"
}
