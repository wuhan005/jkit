package util

import (
	"errors"
	"fmt"
	"io"
	"os"

	"github.com/atotto/clipboard"
)

// ErrNoInput is returned when no data is available from stdin or the clipboard.
var ErrNoInput = errors.New("no input from stdin or clipboard")

// ReadInput reads from stdin first (pipe or redirect) and falls back to the system clipboard.
// The clipboard is only consulted when stdin is a terminal (no data piped in).
func ReadInput() (string, error) {
	if hasStdinData() {
		b, err := io.ReadAll(os.Stdin)
		if err != nil {
			return "", fmt.Errorf("read stdin: %w", err)
		}
		if len(b) > 0 {
			return string(b), nil
		}
	}

	s, err := clipboard.ReadAll()
	if err != nil {
		return "", fmt.Errorf("read clipboard: %w", err)
	}
	if s == "" {
		return "", ErrNoInput
	}
	return s, nil
}

// hasStdinData reports whether stdin is connected to a pipe or redirect rather than a terminal.
// A character device means an interactive terminal with nothing piped in.
func hasStdinData() bool {
	info, err := os.Stdin.Stat()
	if err != nil {
		return false
	}
	return (info.Mode() & os.ModeCharDevice) == 0
}
