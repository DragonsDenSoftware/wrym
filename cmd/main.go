package cmd

import (
	"github.com/syke99/wyvrn-cli/internal/flags"
	"log"
	"os"

	"github.com/syke99/wyvrn-cli/internal/commands"
	"github.com/urfave/cli/v2"
)

var help *bool

func main() {
	app := &cli.App{
		Commands: []*cli.Command{
			commands.New(),
			commands.Run(),
		},
		Flags: []cli.Flag{
			flags.NewFlag(flags.Help, flags.WithDestination(help)),
		},
		Action: func(context *cli.Context) error {
			// TODO: check for help flag and print help data if present
			return nil
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}