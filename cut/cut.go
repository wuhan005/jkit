package cut

import (
	"fmt"
	"strconv"

	jsoniter "github.com/json-iterator/go"
	"github.com/urfave/cli/v2"

	"github.com/wuhan005/jkit/util"
)

func Cmd() *cli.Command {
	return &cli.Command{
		Name:    "cut",
		Usage:   "JSON Cutter",
		Aliases: []string{"c"},
		Action: func(c *cli.Context) error {
			maxDeep, _ := strconv.Atoi(c.Args().First())

			raw := util.ReadClipboard()
			var data interface{}
			err := jsoniter.Unmarshal([]byte(raw), &data)
			if err != nil {
				return err
			}
			fmt.Println(elementParser(data, 0, maxDeep))
			return nil
		},
	}
}

func elementParser(element interface{}, deep int, maxDeep int) string {
	switch data := element.(type) {
	case map[string]interface{}: // Object
		if deep >= maxDeep {
			return fmt.Sprintf("{ %d items dict }", len(data))
		} else if len(data) == 0 {
			return "{}"
		}

		str := "{\n"

		idx := 0
		for k, v := range data {
			idx++
			str += indent(deep)
			str += fmt.Sprintf("    \"%s\": %s", k, elementParser(v, deep+1, maxDeep))
			if idx != len(data) {
				str += ",\n"
			}
		}

		str += "\n" + indent(deep) + "}"
		return str

	case string: // String
		return fmt.Sprintf("\"%s\"", data)

	case int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64: // Number
		return fmt.Sprintf("%d", data)

	case float32, float64:
		return fmt.Sprintf("%f", data)

	case bool: // Boolean
		if data {
			return "true"
		}
		return "false"

	case nil: // NULL
		return "NULL"

	case []interface{}: // Array
		if deep >= maxDeep {
			return fmt.Sprintf("[ %d items array ]", len(data))
		} else if len(data) == 0 {
			return "[]"
		}

		str := "[\n"

		for k, v := range data {
			str += indent(deep+1) + elementParser(v, deep+1, maxDeep)
			if k != len(data)-1 {
				str += ",\n"
			}
		}

		str += "\n" + indent(deep) + "]"
		return str
	}

	return ""
}

func indent(deep int) string {
	str := ""
	for i := 0; i < deep; i++ {
		str += "    "
	}
	return str
}
