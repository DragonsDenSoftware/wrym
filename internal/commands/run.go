package commands

import (
	"fmt"
	"github.com/syke99/wyvrn-cli/internal/constants"
	"github.com/syke99/wyvrn-cli/internal/flags"
	"github.com/urfave/cli/v2"
	"os"
	"os/exec"
	"path/filepath"
)

var runHelp *bool

func cmdRun(ctx *cli.Context) *cli.Command {
	return &cli.Command{
		Name:  constants.RunName,
		Usage: constants.RunUsage,
		Flags: []cli.Flag{
			flags.NewFlag(flags.Dev),
			flags.NewFlag(flags.Staging),
			flags.NewFlag(flags.Prod),
		},
		Action: func(c *cli.Context) error {
			ex, _ := os.Executable()
			curDir := filepath.Dir(ex)

			var cfgLoc string

			if ctx.Bool(constants.DevName) {
				cfgLoc = filepath.Join(curDir, "env", constants.DevName, "config.yaml")
			} else if ctx.Bool(constants.StagingName) {
				cfgLoc = filepath.Join(curDir, "env", constants.StagingName, "config.yaml")
			} else if ctx.Bool(constants.ProdName) {
				cfgLoc = filepath.Join(curDir, "env", constants.ProdName, "config.yaml")
			} else {
				return fmt.Errorf("environment not specified. please use -%s|--%s, -%s|--%s, or -%s|--%s flag",
					constants.DevAlias, constants.DevName, constants.StagingAlias, constants.StagingName, constants.ProdAlias, constants.ProdName)
			}

			return exec.Command("wyvrn", fmt.Sprintf("-cfg=%s", cfgLoc)).Run()
		},
	}
}
