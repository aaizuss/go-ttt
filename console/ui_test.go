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

func TestShow(t *testing.T) {
	var io CommandLine
	var buf bytes.Buffer
	io.Writer = &buf

	messageKey := "tie"
	expectedMessage := "It's a tie!\n"

	io.Show(messageKey)
	result := buf.String()

	if result != expectedMessage {
		t.Errorf("expected output: %s, got %s", expectedMessage, result)
	}
}

func TestGetGameChoiceReturnsChoice(t *testing.T) {
	cli := CliWithInput("1")

	expected := "1"
	result := cli.GetGameChoice()

	if result != expected {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

func TestGetGameChoicePromptsUntilChoiceIsValid(t *testing.T) {
	cli := CliWithInput("abc\n5\n3")

	expected := "3"
	result := cli.GetGameChoice()

	if result != expected {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}
