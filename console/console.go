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
	cli := Console{Writer: os.Stdout, Reader: os.Stdin}
	return &cli
}

func (cli *Console) Read() string {
	var input string

	fmt.Fscanf(cli.Reader, "%s", &input)

	return input
}

func (cli *Console) Write(message string) {
	fmt.Fprintf(cli.Writer, "%v", message)
}
