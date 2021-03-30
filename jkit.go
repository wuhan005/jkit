package main

import (
	"os"

	"github.com/urfave/cli/v2"
	log "unknwon.dev/clog/v2"

	"github.com/wuhan005/jkit/cut"
	"github.com/wuhan005/jkit/format"
	"github.com/wuhan005/jkit/get"
	"github.com/wuhan005/jkit/maker"
)

func main() {
	defer log.Stop()
	err := log.NewConsole()
	if err != nil {
		panic("Failed to init log" + err.Error())
	}

	app := cli.NewApp()
	app.Name = "jkit"
	app.Usage = "JSON CLI Tool"
	app.Commands = []*cli.Command{
		format.Cmd(),
		cut.Cmd(),
		maker.Cmd(),
		get.Cmd(),
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal("Failed to start application: %v", err)
	}
}
