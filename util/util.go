package util

import (
	"github.com/atotto/clipboard"
)

func ReadClipboard() string {
	str, err := clipboard.ReadAll()
	if err != nil {
		return ""
	}
	return str
}
