package cmd

import (
	"github.com/syke99/wyvrn-cli/internal/flags"
	"github.com/syke99/wyvrn-cli/internal/templates"
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
		Flags: []cli.Flag{
			flags.NewFlag(flags.Help, flags.WithDestination(help)),
		},
		Action: func(context *cli.Context) error {
			if *help {
				err := templates.Help(templates.TopLevel)
				return err
			}

			return nil
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
