package flags

import (
	"github.com/DragonsDenSoftware/wrym/internal/pkg/constants"
	"github.com/urfave/cli/v2"
)

func module() *cli.StringFlag {
	return &cli.StringFlag{
		Name:    constants.ModuleName,
		Aliases: []string{constants.ModuleAlias},
		Usage:   constants.ModuleUsage,
	}
}
