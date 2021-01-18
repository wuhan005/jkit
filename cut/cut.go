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
			deep, _ := strconv.Atoi(c.Args().First())

			raw := util.ReadClipboard()
			var data interface{}
			err := jsoniter.Unmarshal([]byte(raw), &data)
			if err != nil {
				return err
			}
			fmt.Println(elementParser(data, deep, deep))
			return nil
		},
	}
}

func elementParser(element interface{}, index int, max int) string {
	switch data := element.(type) {
	case map[string]interface{}: // Object
		if index < 0 {
			return fmt.Sprintf("{ %d items dict }", len(data))
		}

		idx := 0
		str := "{\n"
		for k, v := range data {
			idx++
			for i := index + 1; i < max; i++ {
				str += "    "
			}
			str += fmt.Sprintf("    \"%s\": %s", k, elementParser(v, index-1, max))
			if idx != len(data)-1 {
				str += ","
			}
			str += "\n"
		}
		for i := index + 1; i < max; i++ {
			str += "    "
		}
		str += "}"
		return str

	case string: // String
		return fmt.Sprintf("\"%s\"", data)

	case int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64: // Number
		return fmt.Sprintf("\"%d\"", data)

	case float32, float64:
		return fmt.Sprintf("\"%f\"", data)

	case bool: // Boolean
		if data {
			return "true"
		}
		return "false"

	case nil: // NULL
		return "NULL"

	case []interface{}: // Array
		if index < 0 {
			return fmt.Sprintf("[ %d items array ]", len(data))
		}

		str := "[\n"

		for k, v := range data {
			for i := index; i < max; i++ {
				str += "    "
			}
			str += elementParser(v, index-1, max)
			if k != len(data)-1 {
				str += ","
			}
			str += "\n"
		}

		for i := index + 1; i < max; i++ {
			str += "    "
		}
		str += "]"

		return str
	}

	return ""
}
