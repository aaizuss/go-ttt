package console

import (
	"fmt"
	"io"
	"os"
)

type Console struct {
	Output io.Writer
	Reader io.Reader
}

func New() *Console {
	return &Console{Output: os.Stdout, Reader: os.Stdin}
}

func (cli *Console) Read() string {
	var input string

	fmt.Fscanf(cli.Reader, "%s", &input)

	return input
}

func (cli *Console) Show(message string) {
	fmt.Fprintf(cli.Output, "%v", message)
}
