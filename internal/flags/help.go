package flags

import (
	"github.com/syke99/wyvrn-cli/internal/constants"
	"github.com/urfave/cli/v2"
)

func help() *cli.BoolFlag {
	return &cli.BoolFlag{
		Name:    constants.HelpName,
		Aliases: []string{constants.HelpAlias},
		Usage:   constants.HelpUsage,
	}
}
