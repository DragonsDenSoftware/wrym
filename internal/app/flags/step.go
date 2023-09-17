package flags

import (
	"github.com/syke99/wrym/internal/pkg/constants"
	"github.com/urfave/cli/v2"
)

func step() *cli.StringFlag {
	return &cli.StringFlag{
		Name:    constants.StepName,
		Aliases: []string{constants.StepAlias},
		Usage:   constants.StepUsage,
	}
}
