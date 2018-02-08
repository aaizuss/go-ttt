package console

import (
	"bytes"
	"strings"
	"testing"

	"github.com/aaizuss/go-ttt/board"
)

func cliWithInput(input string) Console {
	var cli Console
	var buf bytes.Buffer
	cli.Output = &buf
	cli.Reader = strings.NewReader(input)
	return cli
}

func TestGetMoveReturnsMove(t *testing.T) {
	board := board.New(3)
	board.MarkSpace(0, "x")
	cli := cliWithInput("5")

	expect := 5
	result := cli.GetMove(board)

	if result != expect {
		t.Errorf("Expected %v, got %v", expect, result)
	}
}

func TestGetMovePromptsForMoveUntilMoveIsValid(t *testing.T) {
	board := board.New(3)
	board.MarkSpace(0, "x")
	cli := cliWithInput("0\n9\n1")

	expect := 1
	result := cli.GetMove(board)

	if result != expect {
		t.Errorf("Expected %v, got %v", expect, result)
	}
}

func TestGetGameChoiceReturnsChoice(t *testing.T) {
	cli := cliWithInput("1")

	expect := "1"
	result := cli.GetGameChoice()

	if result != expect {
		t.Errorf("Expected %v, got %v", expect, result)
	}
}

func TestGetGameChoicePromptsUntilChoiceIsValid(t *testing.T) {
	cli := cliWithInput("abc\n5\n3")

	expect := "3"
	result := cli.GetGameChoice()

	if result != expect {
		t.Errorf("Expected %v, got %v", expect, result)
	}
}

func TestShowOutcome(t *testing.T) {
	var io Console
	var buf bytes.Buffer
	io.Output = &buf
	board := xWinsBoard()

	expectedMessage := "Player x wins!\n"

	io.ShowOutcome(board)
	result := buf.String()

	if result != expectedMessage {
		t.Errorf("expected output: %s, got %s", expectedMessage, result)
	}
}

func xWinsBoard() board.Board {
	board := board.New(3)
	board.MarkSpace(0, "x")
	board.MarkSpace(4, "x")
	board.MarkSpace(8, "x")

	return board
}
