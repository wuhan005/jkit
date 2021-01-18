package main

import (
	"os"

	"github.com/urfave/cli/v2"
	log "unknwon.dev/clog/v2"

	"github.com/wuhan005/jkit/format"
)

func main() {
	err := log.NewConsole()
	if err != nil {
		panic("Failed to init log" + err.Error())
	}

	app := cli.NewApp()
	app.Name = "jkit"
	app.Usage = "JSON CLI Tool"
	app.Commands = []*cli.Command{
		format.Cmd(),
	}
	app.HideHelp = true

	if err := app.Run(os.Args); err != nil {
		log.Fatal("Failed to start application: %v", err)
	}
}
