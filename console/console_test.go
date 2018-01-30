package console

import (
	"bytes"
	"strings"
	"testing"
)

func TestWrite(t *testing.T) {
	var io CommandLine
	var buf bytes.Buffer
	io.Writer = &buf

	message := "Hello, world"

	io.Write(message)
	result := buf.String()

	if result != message {
		t.Errorf("expected output: %s, got %s", message, result)
	}
}

func TestShow(t *testing.T) {
	var io CommandLine
	var buf bytes.Buffer
	io.Writer = &buf

	messageKey := "tie"
	expectedMessage := "It's a tie!\n"

	io.Show(messageKey)
	result := buf.String()

	if result != expectedMessage {
		t.Errorf("expected output: %s, got %s", expectedMessage, result)
	}
}

func TestRead(t *testing.T) {
	var io CommandLine
	input := "Hellothere"

	io.Reader = strings.NewReader(input)

	result := io.Read()

	if result != input {
		t.Fatalf("expected input: %s, got %v", input, result)
	}
}
