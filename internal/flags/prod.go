package flags

import (
	"github.com/syke99/wyvrn-cli/internal/constants"
	"github.com/urfave/cli/v2"
)

func prod() *cli.StringFlag {
	return &cli.StringFlag{
		Name:    constants.ProdName,
		Aliases: []string{constants.ProdAlias},
		Usage:   constants.ProdUsage,
	}
}
