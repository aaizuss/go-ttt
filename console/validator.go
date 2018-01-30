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

// func ValidateChoice(board board.Board, inputMove string) (isValid bool, move int) {
// 	move, err := strconv.Atoi(inputMove)
// 	if err != nil {
// 		return false, -1
// 	}
// 	return board.IsValidMove(move), move
// }
