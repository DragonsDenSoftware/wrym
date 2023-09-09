package flags

import (
	"github.com/syke99/wyvrn-cli/internal/constants"
	"os"
	"path/filepath"

	"github.com/urfave/cli/v2"
)

func directory() *cli.StringFlag {
	ex, err := os.Executable()
	if err != nil {
		panic(err)
	}
	curDir := filepath.Dir(ex)
	return &cli.StringFlag{
		Name:    constants.DirectoryName,
		Aliases: []string{constants.DirectoryAlias},
		Value:   curDir,
		Usage:   constants.DirectoryUsage,
	}
}
