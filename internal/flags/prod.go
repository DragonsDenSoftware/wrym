package flags

import "github.com/urfave/cli/v2"

func prod() *cli.StringFlag {
	return &cli.StringFlag{
		Name:    "prod",
		Aliases: []string{"p"},
		Value:   "",
		Usage:   "specify to run wyvrn with the prod configuration",
	}
}
