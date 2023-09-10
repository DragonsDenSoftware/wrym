package commands

import (
	"github.com/syke99/wyvrn-cli/internal/constants"
	"github.com/syke99/wyvrn-cli/internal/flags"
	"github.com/urfave/cli/v2"
)

var runHelp *bool

func cmdRun() *cli.Command {
	return &cli.Command{
		Name:  constants.RunName,
		Usage: constants.RunUsage,
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
