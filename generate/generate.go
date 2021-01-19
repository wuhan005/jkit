package generate

import (
	"fmt"
	"strings"

	jsoniter "github.com/json-iterator/go"
	"github.com/urfave/cli/v2"

	"github.com/wuhan005/jkit/util"
)

func Cmd() *cli.Command {
	return &cli.Command{
		Name:    "generate",
		Usage:   "JSON Generator",
		Aliases: []string{"g"},
		Action: func(c *cli.Context) error {
			raw := util.ReadClipboard()
			result, err := generator(raw)
			if err != nil {
				return err
			}
			fmt.Println(result)
			return nil
		},
	}
}

func generator(input string) (string, error) {
	group := strings.Split(input, "\n")
	for k, v := range group {
		group[k] = strings.TrimSpace(v)
	}
	json, err := jsoniter.MarshalIndent(group, "", "    ")
	if err != nil {
		return "", err
	}
	return string(json), err
}
