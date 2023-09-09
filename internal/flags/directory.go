package flags

import (
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
		Name:    "dir",
		Aliases: []string{"d"},
		Value:   curDir,
		Usage:   "specify the directory to create your new wyvrn app/module; defaults to the current directory",
	}
}
