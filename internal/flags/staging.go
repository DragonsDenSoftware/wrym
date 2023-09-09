package flags

import "github.com/urfave/cli/v2"

func staging() *cli.StringFlag {
	return &cli.StringFlag{
		Name:    "staging",
		Aliases: []string{"s"},
		Value:   "",
		Usage:   "specify to run wyvrn with the staging configuration",
	}
}
