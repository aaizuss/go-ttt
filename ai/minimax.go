package ai

import (
	"math/rand"
	"time"

	"github.com/aaizuss/tictactoe/board"
)

func ChooseMove(board board.Board, players []string, aiMarker string) int {
	if board.IsEmpty() {
		return randomCorner(corners)
	}
	return minimax(board, 0, players, aiMarker)
}

func score(board board.Board, depth int, aiMarker string) int {
	if board.IsTie() {
		return 0
	} else if aiWin(board, aiMarker) {
		return 100 - depth
	} else {
		return depth - 100
	}
}

func minimax(board board.Board, depth int, players []string, aiMarker string) int {
	if board.GameOver() {
		return score(board, depth, aiMarker)
	}

	scoredMoves := map[int]int{}
	possibleMoves := board.EmptySpaces()

	for _, move := range possibleMoves {
		board.MarkSpace(move, players[0])
		scoredMoves[move] = minimax(board, depth+1, changeTurn(players), aiMarker)
		board.ResetSpace(move)
	}

	bestMove, bestScore := bestPlayerOutcome(scoredMoves, players, aiMarker)
	if depth == 0 {
		return bestMove
	} else {
		return bestScore
	}

}

func bestPlayerOutcome(scoredMoves map[int]int, players []string, aiMarker string) (move int, score int) {
	if isAiTurn(players, aiMarker) {
		return maximizeScore(scoredMoves)
	} else {
		return minimizeScore(scoredMoves)
	}
}

func changeTurn(markers []string) []string {
	return []string{markers[1], markers[0]}
}

func isAiTurn(markers []string, aiMarker string) bool {
	return markers[0] == aiMarker
}

func aiWin(board board.Board, aiMarker string) bool {
	winner, err := board.Winner()
	if err != nil {
		return false
	}
	return winner == aiMarker
}

func maximizeScore(scoredMoves map[int]int) (bestMove int, maxScore int) {
	bestMove, maxScore = worstForMaxPlayer, worstForMaxPlayer
	for move, score := range scoredMoves {
		if score > maxScore {
			maxScore = score
			bestMove = move
		}
	}
	return bestMove, maxScore
}

func minimizeScore(scoredMoves map[int]int) (bestMove int, minScore int) {
	bestMove, minScore = worstForMinPlayer, worstForMinPlayer
	for move, score := range scoredMoves {
		if score < minScore {
			minScore = score
			bestMove = move
		}
	}
	return bestMove, minScore
}

func randomCorner(corners []int) int {
	rand.Seed(time.Now().UnixNano())
	return corners[rand.Intn(len(corners)-1)]
}

const (
	worstForMaxPlayer = -100
	worstForMinPlayer = 100
)

var corners = []int{0, 2, 6, 8}
