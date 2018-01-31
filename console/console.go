package console

import (
	"fmt"
	"io"
	"os"
)

type CommandLine struct {
	Writer io.Writer
	Reader io.Reader
}

func New() *CommandLine {
	cli := CommandLine{Writer: os.Stdout, Reader: os.Stdin}
	return &cli
}

func (cli *CommandLine) Read() string {
	var input string

	fmt.Fscanf(cli.Reader, "%s", &input)

	return input
}

func (cli *CommandLine) Write(message string) {
	fmt.Fprintf(cli.Writer, "%v", message)
}
