package commands

import (
	"fmt"
	"github.com/syke99/wyvrn-cli/internal/constants"
	"github.com/syke99/wyvrn-cli/internal/flags"
	"github.com/urfave/cli/v2"
	"os/exec"
	"path/filepath"
)

var runEnv *string

func cmdRun(dir string) *cli.Command {
	return &cli.Command{
		Name:  constants.RunName,
		Usage: constants.RunUsage,
		Flags: []cli.Flag{
			flags.NewFlag(flags.Env, flags.WithDestination(runEnv), flags.Required(constants.RunName)),
		},
		Action: func(ctx *cli.Context) error {
			env := ctx.String(constants.EnvName)

			cfgLoc := filepath.Join(dir, "env", env, "config.yaml")

			return exec.Command("wyvrn", fmt.Sprintf("-cfg=%s", cfgLoc)).Run()
		},
	}
}
