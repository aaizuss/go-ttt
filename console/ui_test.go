package console

import (
	"bytes"
	"strings"
	"testing"

	"github.com/aaizuss/tictactoe/board"
)

func CliWithInput(input string) Console {
	var cli Console
	var buf bytes.Buffer
	cli.Output = &buf
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
	var io Console
	var buf bytes.Buffer
	io.Output = &buf

	message := "It's a tie!\n"

	io.Show(message)
	result := buf.String()

	if result != message {
		t.Errorf("expected output: %s, got %s", message, result)
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
