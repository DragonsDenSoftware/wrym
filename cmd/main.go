package main

import (
	"github.com/syke99/wyvrn-cli/internal/commands"
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

var help *bool

func main() {
	app := &cli.App{}

	cCtx := cli.NewContext(app, nil, nil)

	app.Commands = []*cli.Command{
		commands.NewCommand(cCtx, commands.New),
		commands.NewCommand(cCtx, commands.Run),
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
