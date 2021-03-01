package maker

import (
	"fmt"
	"strings"

	jsoniter "github.com/json-iterator/go"
	"github.com/urfave/cli/v2"

	"github.com/wuhan005/jkit/util"
)

func Cmd() *cli.Command {
	return &cli.Command{
		Name:    "maker",
		Usage:   "JSON Maker",
		Aliases: []string{"m"},
		Flags: []cli.Flag{
			&cli.BoolFlag{Name: "repeat", Aliases: []string{"r"}},
		},
		Action: func(c *cli.Context) error {
			raw := util.ReadClipboard()
			repeat := c.Bool("repeat")
			result, err := generator(raw, repeat)
			if err != nil {
				return err
			}
			fmt.Println(result)
			return nil
		},
	}
}

func generator(input string, removeRepeat bool) (string, error) {
	group := strings.Split(input, "\n")

	var newGroup []string
	dict := map[string]struct{}{}

	for _, v := range group {
		if _, ok := dict[v]; removeRepeat && ok {
			continue
		}
		dict[v] = struct{}{}

		newGroup = append(newGroup, v)
	}

	json, err := jsoniter.MarshalIndent(newGroup, "", "    ")
	if err != nil {
		return "", err
	}
	return string(json), err
}
