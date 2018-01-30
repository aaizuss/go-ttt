package console

import (
	"bytes"
	"github.com/aaizuss/tictactoe/board"
	"strings"
	"testing"
)

func CliWithInput(input string) CommandLine {
	var cli CommandLine
	var buf bytes.Buffer
	cli.Writer = &buf
	cli.Reader = strings.NewReader(input)
	return cli
}

func TestGetMoveReturnsMove(t *testing.T) {
	board := board.New(3)
	board.MarkSpace(0, "x")
	cli := CliWithInput("5")

	expected := 5
	result := cli.GetMove(board)

	if result != expected {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

func TestGetMovePromptsForMoveUntilMoveIsValid(t *testing.T) {
	board := board.New(3)
	board.MarkSpace(0, "x")
	cli := CliWithInput("0\n9\n1")

	expected := 1
	result := cli.GetMove(board)

	if result != expected {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}
