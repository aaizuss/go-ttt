package console

import (
	"fmt"
	"io"
	"os"
)

type Console struct {
	Writer io.Writer
	Reader io.Reader
}

func New() *Console {
	return &Console{Writer: os.Stdout, Reader: os.Stdin}
}

func (cli *Console) Read() string {
	var input string

	fmt.Fscanf(cli.Reader, "%s", &input)

	return input
}

func (cli *Console) Write(message string) {
	fmt.Fprintf(cli.Writer, "%v", message)
}
