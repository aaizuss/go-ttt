package game

import (
	"github.com/aaizuss/tictactoe/board"
	"github.com/aaizuss/tictactoe/console"

	"bytes"
	"reflect"
	"strings"

	"testing"
)

func CliWithInput(input string) console.CommandLine {
	var cli console.CommandLine
	var buf bytes.Buffer
	cli.Writer = &buf
	cli.Reader = strings.NewReader(input)
	return cli
}

func TestTogglePlayer(t *testing.T) {
	cli := CliWithInput("") // going to need to refactor CommandLine so I can mock it
	game := Game{board: board.New(3), players: []string{"x", "o"}, ui: cli}

	expected := []string{"o", "x"}
	game.TogglePlayer()
	result := game.players

	if !reflect.DeepEqual(expected, result) {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}
