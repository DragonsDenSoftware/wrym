package flags

import (
	"github.com/syke99/wyvrn-cli/internal/constants"
	"github.com/urfave/cli/v2"
)

func language() *cli.StringFlag {
	return &cli.StringFlag{
		Name:    constants.LanguageName,
		Aliases: []string{constants.LanguageAlias},
		Usage:   constants.LanguageUsage,
	}
}
