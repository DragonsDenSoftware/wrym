package flags

import (
	"github.com/syke99/wyvrn-cli/internal/pkg/constants"
	"github.com/urfave/cli/v2"
)

func name() *cli.StringFlag {
	return &cli.StringFlag{
		Name:    constants.NameName,
		Aliases: []string{constants.NameAlias},
		Usage:   constants.NameUsage,
	}
}
