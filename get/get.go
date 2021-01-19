package get

import (
	"fmt"
	"strconv"
	"strings"

	jsoniter "github.com/json-iterator/go"
	"github.com/urfave/cli/v2"
	log "unknwon.dev/clog/v2"

	"github.com/wuhan005/jkit/util"
)

func Cmd() *cli.Command {
	return &cli.Command{
		Name:    "get",
		Usage:   "JSON Getter",
		Aliases: []string{"g"},
		Action: func(c *cli.Context) error {
			raw := util.ReadClipboard()
			chain := c.Args().First()
			result, err := get(raw, chain)
			if err != nil {
				return err
			}
			fmt.Println(result)
			return nil
		},
	}
}

func get(raw, chain string) (string, error) {
	var json interface{}
	err := jsoniter.Unmarshal([]byte(raw), &json)
	if err != nil {
		return "", err
	}

	chainGroups := strings.Split(strings.TrimSpace(chain), ".")
	return parser(json, chainGroups, 0), nil
}

func parser(element interface{}, groups []string, index int) string {
	if index >= len(groups) {
		json, _ := jsoniter.MarshalIndent(element, "", "    ")
		return string(json)
	}

	switch data := element.(type) {
	case map[string]interface{}: // Object
		v, ok := data[groups[index]]
		if !ok {
			log.Fatal("%v not found", groups[index])
		}
		return parser(v, groups, index+1)
	case []interface{}: // Array
		idx, err := strconv.Atoi(groups[index])
		if err != nil {
			log.Fatal("%v must be a number", groups[index])
		}
		if idx >= len(data) {
			log.Fatal("Index %v out of range, max index is %d", groups[index], len(data)-1)
		}
		return parser(data[idx], groups, index+1)
	case string: // String
		return data
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
	}

	return ""
}
