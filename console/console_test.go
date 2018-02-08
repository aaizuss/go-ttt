package console

import (
	"bytes"
	"strings"
	"testing"
)

func TestShow(t *testing.T) {
	var io Console
	var buf bytes.Buffer
	io.Output = &buf

	message := "Hello, world"

	io.Show(message)
	output := buf.String()

	if output != message {
		t.Errorf("expected output: %s, got: %s", message, output)
	}
}

func TestRead(t *testing.T) {
	var io Console
	expect := "Hello"

	io.Reader = strings.NewReader(expect)

	result := io.Read()

	if result != expect {
		t.Fatalf("expected input: %s, got %s", expect, result)
	}
}
