package flags

import "github.com/urfave/cli/v2"

func dev() *cli.StringFlag {
	return &cli.StringFlag{
		Name:    "dev",
		Aliases: []string{"d"},
		Value:   "",
		Usage:   "specify to run wyvrn with the dev configuration",
	}
}
