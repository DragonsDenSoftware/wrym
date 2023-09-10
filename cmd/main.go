package main

import (
	"log"
	"os"

	"github.com/syke99/wyvrn-cli/internal/commands"
	"github.com/urfave/cli/v2"
)

var help *bool

func main() {
	app := &cli.App{
		Commands: []*cli.Command{
			commands.NewCommand(commands.New),
			commands.NewCommand(commands.Run),
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
