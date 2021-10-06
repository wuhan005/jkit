package format

import (
	"encoding/json"
	"fmt"

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
			err := json.Unmarshal([]byte(raw), &data)
			if err != nil {
				return err
			}

			jsonBytes, err := json.MarshalIndent(data, "", "    ")
			if err != nil {
				return err
			}

			fmt.Println(string(jsonBytes))
			return nil
		},
	}
}
