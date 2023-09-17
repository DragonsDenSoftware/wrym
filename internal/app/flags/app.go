package flags

import (
	"github.com/syke99/wrym/internal/pkg/constants"
	"github.com/urfave/cli/v2"
)

func app() *cli.BoolFlag {
	return &cli.BoolFlag{
		Name:    constants.ConfigName,
		Aliases: []string{constants.ConfigAlias},
		Usage:   constants.ConfigUsage,
	}
}