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
	output := buf.String()

	if output != message {
		t.Errorf("expected output: %s, got %s", message, output)
	}
}

func TestRead(t *testing.T) {
	var io Console
	want := "Hello"

	io.Reader = strings.NewReader(want)

	got := io.Read()

	if got != want {
		t.Fatalf("expected input: %s, got %v", want, got)
	}
}
