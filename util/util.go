package util

import (
	"io"
	"os"

	"github.com/atotto/clipboard"
	"github.com/pkg/errors"
)

func ReadInput() string {
	input, err := readFromStdin()
	if err == nil {
		return input
	}

	return readClipboard()
}

func readFromStdin() (string, error) {
	stdinInfo, err := os.Stdin.Stat()
	if err != nil {
		return "", errors.Wrap(err, "stat stdin")
	}

	if stdinInfo.Size() == 0 {
		return "", errors.New("empty size")
	}

	input, err := io.ReadAll(os.Stdin)
	return string(input), err
}

func readClipboard() string {
	str, err := clipboard.ReadAll()
	if err != nil {
		return ""
	}
	return str
}
