package commands

import (
	"fmt"
	"github.com/DragonsDenSoftware/wrym/internal/app/flags"
	"github.com/DragonsDenSoftware/wrym/internal/pkg/constants"
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
			flags.NewFlag(flags.Env, flags.WithDestination(runEnv), flags.Required()),
		},
		Action: func(ctx *cli.Context) error {
			env := ctx.String(constants.EnvName)

			cfgLoc := filepath.Join(dir, "env", env, "config.yaml")

			return exec.Command("wyvrn", fmt.Sprintf("-cfg=%s", cfgLoc)).Run()
		},
	}
}
