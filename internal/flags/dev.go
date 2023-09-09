package flags

import (
	"github.com/syke99/wyvrn-cli/internal/constants"
	"github.com/urfave/cli/v2"
)

func dev() *cli.StringFlag {
	return &cli.StringFlag{
		Name:    constants.DevName,
		Aliases: []string{constants.DevAlias},
		Usage:   constants.DevUsage,
	}
}
