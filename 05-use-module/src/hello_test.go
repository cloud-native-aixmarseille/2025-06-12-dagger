package main

import (
	"io"
	"os"
	"testing"
)

func TestMainFunctionWithArgument(t *testing.T) {
	oldArgs := os.Args
	oldStdout := os.Stdout
	defer func() {
		os.Args = oldArgs
		os.Stdout = oldStdout
	}()

	r, w, _ := os.Pipe()
	os.Stdout = w

	os.Args = []string{"cmd", "World"}
	main()

	w.Close()
	out, _ := io.ReadAll(r)
	os.Stdout = oldStdout // Restore stdout before logging
	os.Args = oldArgs

	expected := "Hello World 👋\n"
	if string(out) != expected {
		t.Errorf("Expected output %q, but got %q", expected, string(out))
	}
}

func TestMainFunctionWithoutArgument(t *testing.T) {
	oldArgs := os.Args
	oldStdout := os.Stdout
	defer func() {
		os.Args = oldArgs
		os.Stdout = oldStdout
	}()

	r, w, _ := os.Pipe()
	os.Stdout = w

	os.Args = []string{"cmd"}
	main()

	w.Close()
	out, _ := io.ReadAll(r)
	os.Stdout = oldStdout // Restore stdout before logging
	os.Args = oldArgs

	expected := "Usage: hello <something>\n"
	if string(out) != expected {
		t.Errorf("Expected output %q, but got %q", expected, string(out))
	}
}

func TestMainFunctionWithMultipleArguments(t *testing.T) {
	oldArgs := os.Args
	oldStdout := os.Stdout
	defer func() {
		os.Args = oldArgs
		os.Stdout = oldStdout
	}()

	r, w, _ := os.Pipe()
	os.Stdout = w

	os.Args = []string{"cmd", "First", "Second"}
	main()

	w.Close()
	out, _ := io.ReadAll(r)
	os.Stdout = oldStdout
	os.Args = oldArgs

	// The program currently only uses the first argument after the command name.
	expected := "Hello First 👋\n"
	if string(out) != expected {
		t.Errorf("Expected output %q, but got %q", expected, string(out))
	}
}

func TestMainFunctionWithEmptyArgument(t *testing.T) {
	oldArgs := os.Args
	oldStdout := os.Stdout
	defer func() {
		os.Args = oldArgs
		os.Stdout = oldStdout
	}()

	r, w, _ := os.Pipe()
	os.Stdout = w

	os.Args = []string{"cmd", ""}
	main()

	w.Close()
	out, _ := io.ReadAll(r)
	os.Stdout = oldStdout
	os.Args = oldArgs

	expected := "Hello  👋\n"
	if string(out) != expected {
		t.Errorf("Expected output %q, but got %q", expected, string(out))
	}
}
