package get

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"

	"github.com/urfave/cli/v2"

	"github.com/wuhan005/jkit/internal/util"
)

// Cmd extracts a JSON sub-node by an `a.b.0.c` style dotted path.
// Arrays use decimal indices; string leaves print raw (no quotes), other leaves print as JSON literals.
func Cmd() *cli.Command {
	return &cli.Command{
		Name:        "get",
		Aliases:     []string{"g"},
		Usage:       "Extract a JSON sub-element by path",
		ArgsUsage:   "<path>",
		Description: "Read JSON from stdin or the system clipboard and extract the value at <path>, e.g. result.items.0.title.",
		Action: func(c *cli.Context) error {
			raw, err := util.ReadInput()
			if err != nil {
				return err
			}
			path := c.Args().First()
			out, err := Get([]byte(raw), path)
			if err != nil {
				return err
			}
			fmt.Println(out)
			return nil
		},
	}
}

// Get parses raw and walks the dotted path; an empty path returns the whole document formatted.
func Get(raw []byte, path string) (string, error) {
	dec := json.NewDecoder(bytes.NewReader(raw))
	dec.UseNumber()

	var data any
	if err := dec.Decode(&data); err != nil {
		return "", fmt.Errorf("parse json: %w", err)
	}

	parts := splitPath(path)
	node, err := walk(data, parts)
	if err != nil {
		return "", err
	}
	return format(node)
}

func splitPath(path string) []string {
	path = strings.TrimSpace(path)
	if path == "" {
		return nil
	}
	return strings.Split(path, ".")
}

func walk(node any, parts []string) (any, error) {
	for i, p := range parts {
		switch x := node.(type) {
		case map[string]any:
			v, ok := x[p]
			if !ok {
				return nil, fmt.Errorf("key %q not found at %s", p, pathOf(parts, i))
			}
			node = v
		case []any:
			idx, err := strconv.Atoi(p)
			if err != nil {
				return nil, fmt.Errorf("array index must be integer at %s, got %q", pathOf(parts, i), p)
			}
			if idx < 0 || idx >= len(x) {
				return nil, fmt.Errorf("index %d out of range [0, %d) at %s", idx, len(x), pathOf(parts, i))
			}
			node = x[idx]
		default:
			return nil, fmt.Errorf("cannot descend into %T at %s", x, pathOf(parts, i))
		}
	}
	return node, nil
}

func pathOf(parts []string, i int) string {
	return strings.Join(parts[:i+1], ".")
}

// format prints string leaves as raw strings (friendly for shell pipes); everything else as a JSON literal.
func format(node any) (string, error) {
	if s, ok := node.(string); ok {
		return s, nil
	}
	out, err := json.MarshalIndent(node, "", "    ")
	if err != nil {
		return "", fmt.Errorf("marshal json: %w", err)
	}
	return string(out), nil
}
