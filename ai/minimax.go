package ai

import (
	"github.com/aaizuss/tictactoe/board"
)

func ChooseMove(board board.Board, players []string, aiMarker string) int {
	return minimax(board, 0, players, aiMarker)
}

func changeTurn(markers []string) []string {
	return []string{markers[1], markers[0]}
}

func isAiTurn(markers []string, aiMarker string) bool {
	return markers[0] == aiMarker
}

func aiWin(board board.Board, aiMarker string) bool {
	winner, _ := board.Winner()
	return winner == aiMarker
}

// can only be called when the game is over!
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

func bestPlayerOutcome(scoredMoves map[int]int, players []string, aiMarker string) (int, int) {
	if isAiTurn(players, aiMarker) {
		return maximizeScore(scoredMoves)
	} else {
		return minimizeScore(scoredMoves)
	}
}

func maximizeScore(scoredMoves map[int]int) (bestMove int, maxScore int) {
	bestMove, maxScore = -100, -100
	for move, score := range scoredMoves {
		// fmt.Printf("current bestmove: %v, score: %v\n", bestMove, maxScore)
		// fmt.Printf("looking at move: %v, score: %v\n", move, score)
		if score > maxScore {
			//fmt.Printf("changing to move %v with max score: %v\n", bestMove, maxScore)
			maxScore = score
			bestMove = move
		}
	}
	return bestMove, maxScore
}

func minimizeScore(scoredMoves map[int]int) (bestMove int, minScore int) {
	bestMove, minScore = 100, 100
	for move, score := range scoredMoves {
		if score < minScore {
			minScore = score
			bestMove = move
		}
	}
	return bestMove, minScore
}
