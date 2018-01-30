package console

import (
	"fmt"
	"io"
	"os"
)

type IO interface {
	Read() string
	Write(message string)
}

type CommandLine struct {
	Writer io.Writer
	Reader io.Reader
}

func New() CommandLine {
	return CommandLine{Writer: os.Stdout, Reader: os.Stdin}
}

func (cli CommandLine) Read() string {
	var input string

	fmt.Fscanf(cli.Reader, "%s", &input)

	return input
}

func (cli CommandLine) Write(output string) {
	fmt.Fprintf(cli.Writer, "%v", output)
}

func (cli CommandLine) Show(key string) {
	cli.Write(messages[key])
}

// put in a json file at some point?
var messages = map[string]string{
	"welcome":      "|----------------------------|\n|-- Welcome to Tic Tac Toe --|\n|----------------------------|\n",
	"tie":          "It's a tie!\n",
	"choose-space": "Enter a number 0-8 to mark that position on the board: ",
	"invalid-move": "You can't move there. ",
	"taken-space":  "That space is taken. ",
}
