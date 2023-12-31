package flags

import (
	"github.com/DragonsDenSoftware/wrym/internal/pkg/constants"
	"github.com/urfave/cli/v2"
)

func config() *cli.BoolFlag {
	return &cli.BoolFlag{
		Name:    constants.ConfigName,
		Aliases: []string{constants.ConfigAlias},
		Usage:   constants.ConfigUsage,
	}
}
