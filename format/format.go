package format

import (
	"fmt"

	jsoniter "github.com/json-iterator/go"
	"github.com/urfave/cli/v2"

	"github.com/wuhan005/jkit/util"
)

func Cmd() *cli.Command {
	return &cli.Command{
		Name:    "format",
		Usage:   "JSON Format",
		Aliases: []string{"f"},
		Action: func(c *cli.Context) error {
			raw := util.ReadInput()
			var data interface{}
			err := jsoniter.Unmarshal([]byte(raw), &data)
			if err != nil {
				return err
			}

			json, err := jsoniter.MarshalIndent(data, "", "    ")
			if err != nil {
				return err
			}

			fmt.Println(string(json))
			return nil
		},
	}
}
