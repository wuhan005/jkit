package format

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/urfave/cli/v2"

	"github.com/wuhan005/jkit/internal/util"
)

// Cmd pretty-prints JSON from stdin or the clipboard with 4-space indent, preserving numeric precision.
func Cmd() *cli.Command {
	return &cli.Command{
		Name:        "format",
		Aliases:     []string{"f"},
		Usage:       "Pretty-print JSON",
		Description: "Read JSON from stdin or the system clipboard and print it indented with 4 spaces.",
		Action: func(c *cli.Context) error {
			raw, err := util.ReadInput()
			if err != nil {
				return err
			}
			out, err := Format([]byte(raw))
			if err != nil {
				return err
			}
			fmt.Println(out)
			return nil
		},
	}
}

// Format parses raw and returns a 4-space indented string. UseNumber preserves the original numeric literal and avoids float precision loss.
func Format(raw []byte) (string, error) {
	dec := json.NewDecoder(bytes.NewReader(raw))
	dec.UseNumber()

	var data any
	if err := dec.Decode(&data); err != nil {
		return "", fmt.Errorf("parse json: %w", err)
	}

	out, err := json.MarshalIndent(data, "", "    ")
	if err != nil {
		return "", fmt.Errorf("marshal json: %w", err)
	}
	return string(out), nil
}
