package flags

import (
	"github.com/syke99/wyvrn-cli/internal/constants"
	"github.com/urfave/cli/v2"
)

func env() *cli.StringFlag {
	return &cli.StringFlag{
		Name:    constants.EnvName,
		Aliases: []string{constants.EnvAlias},
		Usage:   constants.EnvUsage,
	}
}
