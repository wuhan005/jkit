package cut

import (
	"bytes"
	"encoding/json"
	"fmt"
	"math"
	"sort"
	"strconv"
	"strings"

	"github.com/urfave/cli/v2"

	"github.com/wuhan005/jkit/internal/util"
)

const indentUnit = "    "

// Cmd folds JSON: nodes deeper than maxDepth are summarized as `{ N items dict }` or `[ N items array ]`.
// When no depth argument is given, nothing is folded (behaves like format but with sorted keys).
func Cmd() *cli.Command {
	return &cli.Command{
		Name:        "cut",
		Aliases:     []string{"c"},
		Usage:       "Fold JSON to a maximum depth",
		ArgsUsage:   "<depth>",
		Description: "Read JSON from stdin or the system clipboard and collapse nodes deeper than <depth>. Omit <depth> to keep the full structure.",
		Action: func(c *cli.Context) error {
			maxDepth := math.MaxInt
			if arg := c.Args().First(); arg != "" {
				n, err := strconv.Atoi(arg)
				if err != nil {
					return fmt.Errorf("invalid depth %q: %w", arg, err)
				}
				if n < 0 {
					return fmt.Errorf("depth must be non-negative, got %d", n)
				}
				maxDepth = n
			}

			raw, err := util.ReadInput()
			if err != nil {
				return err
			}

			out, err := Cut([]byte(raw), maxDepth)
			if err != nil {
				return err
			}
			fmt.Println(out)
			return nil
		},
	}
}

// Cut parses raw and folds it at maxDepth. Object keys are emitted in lexicographic order for stable output.
func Cut(raw []byte, maxDepth int) (string, error) {
	dec := json.NewDecoder(bytes.NewReader(raw))
	dec.UseNumber()

	var data any
	if err := dec.Decode(&data); err != nil {
		return "", fmt.Errorf("parse json: %w", err)
	}

	var b strings.Builder
	render(&b, data, 0, maxDepth)
	return b.String(), nil
}

func render(b *strings.Builder, v any, depth, maxDepth int) {
	switch x := v.(type) {
	case map[string]any:
		renderObject(b, x, depth, maxDepth)
	case []any:
		renderArray(b, x, depth, maxDepth)
	case string:
		b.WriteString(strconv.Quote(x))
	case json.Number:
		b.WriteString(x.String())
	case bool:
		if x {
			b.WriteString("true")
		} else {
			b.WriteString("false")
		}
	case nil:
		b.WriteString("null")
	default:
		// Unexpected type (UseNumber should keep us in the cases above); fall back to generic marshal.
		raw, err := json.Marshal(x)
		if err != nil {
			b.WriteString("null")
			return
		}
		b.Write(raw)
	}
}

func renderObject(b *strings.Builder, m map[string]any, depth, maxDepth int) {
	if len(m) == 0 {
		b.WriteString("{}")
		return
	}
	if depth >= maxDepth {
		fmt.Fprintf(b, "{ %d items dict }", len(m))
		return
	}

	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	b.WriteString("{\n")
	inner := strings.Repeat(indentUnit, depth+1)
	for i, k := range keys {
		b.WriteString(inner)
		b.WriteString(strconv.Quote(k))
		b.WriteString(": ")
		render(b, m[k], depth+1, maxDepth)
		if i != len(keys)-1 {
			b.WriteString(",")
		}
		b.WriteString("\n")
	}
	b.WriteString(strings.Repeat(indentUnit, depth))
	b.WriteString("}")
}

func renderArray(b *strings.Builder, a []any, depth, maxDepth int) {
	if len(a) == 0 {
		b.WriteString("[]")
		return
	}
	if depth >= maxDepth {
		fmt.Fprintf(b, "[ %d items array ]", len(a))
		return
	}

	b.WriteString("[\n")
	inner := strings.Repeat(indentUnit, depth+1)
	for i, item := range a {
		b.WriteString(inner)
		render(b, item, depth+1, maxDepth)
		if i != len(a)-1 {
			b.WriteString(",")
		}
		b.WriteString("\n")
	}
	b.WriteString(strings.Repeat(indentUnit, depth))
	b.WriteString("]")
}
