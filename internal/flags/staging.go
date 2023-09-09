package flags

import (
	"github.com/syke99/wyvrn-cli/internal/constants"
	"github.com/urfave/cli/v2"
)

func staging() *cli.StringFlag {
	return &cli.StringFlag{
		Name:    constants.StagingName,
		Aliases: []string{constants.StagingAlias},
		Usage:   constants.StagingUsage,
	}
}
