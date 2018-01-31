package console

import (
	"fmt"
	"io"
	"os"
)

type ReadWriter interface {
	Read() string
	Write(message string)
}

type CommandLine struct {
	Writer io.Writer
	Reader io.Reader
}

type UIReadWriter interface {
	ReadWriter
	UI
}

func New() CommandLine {
	return CommandLine{Writer: os.Stdout, Reader: os.Stdin}
}

func (cli CommandLine) Read() string {
	var input string

	fmt.Fscanf(cli.Reader, "%s", &input)

	return input
}

func (cli CommandLine) Write(message string) {
	fmt.Fprintf(cli.Writer, "%v", message)
}
