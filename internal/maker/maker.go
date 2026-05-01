package maker

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/urfave/cli/v2"

	"github.com/wuhan005/jkit/internal/util"
)

// Cmd splits stdin or the clipboard into lines and emits a JSON string array. Each line is trimmed; empty lines are always dropped.
func Cmd() *cli.Command {
	return &cli.Command{
		Name:        "maker",
		Aliases:     []string{"m"},
		Usage:       "Build a JSON string array from line-separated input",
		Description: "Read text from stdin or the system clipboard and emit a JSON array, one element per non-empty line.",
		Flags: []cli.Flag{
			&cli.BoolFlag{Name: "unique", Aliases: []string{"u", "r"}, Usage: "drop duplicate lines"},
		},
		Action: func(c *cli.Context) error {
			raw, err := util.ReadInput()
			if err != nil {
				return err
			}
			out, err := Make(raw, c.Bool("unique"))
			if err != nil {
				return err
			}
			fmt.Println(out)
			return nil
		},
	}
}

// Make splits on \n (also handling \r\n), trims each line, and drops empties. When unique is true, duplicates are removed preserving first-seen order.
func Make(input string, unique bool) (string, error) {
	lines := strings.Split(strings.ReplaceAll(input, "\r\n", "\n"), "\n")

	out := make([]string, 0, len(lines))
	seen := make(map[string]struct{}, len(lines))
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}
		if unique {
			if _, ok := seen[line]; ok {
				continue
			}
			seen[line] = struct{}{}
		}
		out = append(out, line)
	}

	b, err := json.MarshalIndent(out, "", "    ")
	if err != nil {
		return "", fmt.Errorf("marshal json: %w", err)
	}
	return string(b), nil
}
