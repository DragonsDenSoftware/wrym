package commands

import (
	"github.com/syke99/wyvrn-cli/internal/flags"
	"github.com/urfave/cli/v2"
)

func Run() *cli.Command {
	return &cli.Command{
		Name:  "new",
		Usage: "generate a new wyvrn app/module",
		Flags: []cli.Flag{
			flags.NewFlag(flags.Dev),
			flags.NewFlag(flags.Staging),
			flags.NewFlag(flags.Prod),
		},
		Action: func(c *cli.Context) error {
			return runWyvrn()
		},
	}
}

func runWyvrn() error {
	// TODO: body
	return nil
}
