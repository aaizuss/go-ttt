package ai

import (
	"reflect"
	"testing"

	"github.com/aaizuss/tictactoe/board"
)

var (
	aiMarker     = "x"
	aiTurn       = []string{aiMarker, "o"}
	opponentTurn = []string{"o", aiMarker}
)

func TestMinimaxReturnsWinningMove(t *testing.T) {
	board := xWinsOn8()
	expect := 8

	result := minimax(board, 0, aiTurn, aiMarker)

	if result != expect {
		t.Errorf("Expected %v, got %v", expect, result)
	}
}

func TestMinimaxBlocksOpponent(t *testing.T) {
	board := xMustBlockOn7()
	expect := 7

	result := minimax(board, 0, aiTurn, aiMarker)

	if result != expect {
		t.Errorf("Expected %v, got %v", expect, result)
	}
}

func TestChoosesCornerWhenFirstPlayerOnEmptyBoard(t *testing.T) {
	board := board.New(3)
	corners := []int{0, 2, 6, 8}

	move := ChooseMove(board, aiTurn, aiMarker)

	if !contains(corners, move) {
		t.Errorf("Expected move to be a corner (%v), got %v", corners, move)
	}
}

func TestChangeTurn(t *testing.T) {
	markers := []string{"x", "o"}
	expect := []string{"o", "x"}

	result := changeTurn(markers)

	if !reflect.DeepEqual(expect, result) {
		t.Errorf("Expected %v, got %v", expect, result)
	}
}

func TestIsAiTurnReturnsTrue(t *testing.T) {
	expect := true
	result := isAiTurn(aiTurn, aiMarker)

	if result != expect {
		t.Errorf("Expected %v, got %v", expect, result)
	}
}

func TestIsAiTurnReturnsFalse(t *testing.T) {
	markers := []string{"o", "x"}
	aiMarker := "x"

	expect := false
	result := isAiTurn(markers, aiMarker)

	if result != expect {
		t.Errorf("Expected %v, got %v", expect, result)
	}
}

func TestScoreReturnsZeroForTie(t *testing.T) {
	board := tieBoard()

	expect := 0
	result := score(board, 0, aiMarker)

	if result != expect {
		t.Errorf("Expected %v, got %v", expect, result)
	}
}

func TestBestPlayerOutcomesDoesMaxCalcWhenAiTurn(t *testing.T) {
	scoredMoves := map[int]int{
		6: -98, 7: 0, 8: -98, 3: -98, 5: -98,
	}

	expectedMove, expectedScore := 7, 0
	resultMove, resultScore := bestPlayerOutcome(scoredMoves, aiTurn, aiMarker)

	if resultMove != expectedMove || resultScore != expectedScore {
		t.Errorf("Expected move:%v, score:%v but got move:%v, score:%v", expectedMove, expectedScore, resultMove, resultScore)
	}
}

func TestBestPlayerOutcomesDoesMinCalcWhenOpponentTurn(t *testing.T) {
	scoredMoves := map[int]int{
		2: -10,
		8: 35,
		4: -12,
		3: 2,
	}

	expectedMove, expectedScore := 4, -12
	resultMove, resultScore := bestPlayerOutcome(scoredMoves, opponentTurn, aiMarker)

	if resultMove != expectedMove || resultScore != expectedScore {
		t.Errorf("Expected move:%v, score:%v but got move:%v, score:%v", expectedMove, expectedScore, resultMove, resultScore)
	}
}

func xWinsOn8() board.Board {
	board := board.New(3)
	board.MarkSpace(0, "x")
	board.MarkSpace(1, "o")
	board.MarkSpace(4, "x")
	board.MarkSpace(2, "o")

	return board
}

func xMustBlockOn7() board.Board {
	board := board.New(3)
	board.MarkSpace(0, "x")
	board.MarkSpace(1, "o")
	board.MarkSpace(2, "x")
	board.MarkSpace(4, "o")

	return board
}

func tieBoard() board.Board {
	board := board.New(3)
	board.MarkSpace(0, "x")
	board.MarkSpace(1, "x")
	board.MarkSpace(2, "o")
	board.MarkSpace(3, "o")
	board.MarkSpace(4, "o")
	board.MarkSpace(5, "x")
	board.MarkSpace(6, "x")
	board.MarkSpace(7, "o")
	board.MarkSpace(8, "x")

	return board
}

func contains(arr []int, element int) bool {
	for _, item := range arr {
		if item == element {
			return true
		}
	}
	return false
}
