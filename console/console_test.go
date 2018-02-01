package console

import (
	"bytes"
	"strings"
	"testing"
)

func TestWrite(t *testing.T) {
	var io Console
	var buf bytes.Buffer
	io.Writer = &buf

	message := "Hello, world"

	io.Write(message)
	result := buf.String()

	if result != message {
		t.Errorf("expected output: %s, got %s", message, result)
	}
}

func TestRead(t *testing.T) {
	var io Console
	input := "Hellothere"

	io.Reader = strings.NewReader(input)

	result := io.Read()

	if result != input {
		t.Fatalf("expected input: %s, got %v", input, result)
	}
}
