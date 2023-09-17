package flags

import (
	"github.com/DragonsDenSoftware/wrym/internal/pkg/constants"
	"github.com/urfave/cli/v2"
)

func language() *cli.StringFlag {
	return &cli.StringFlag{
		Name:    constants.LanguageName,
		Aliases: []string{constants.LanguageAlias},
		Usage:   constants.LanguageUsage,
	}
}
