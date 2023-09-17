package commands

import (
	"fmt"
	"github.com/syke99/wyvrn-cli/internal/app/flags"
	constants2 "github.com/syke99/wyvrn-cli/internal/pkg/constants"
	"github.com/urfave/cli/v2"
	"os/exec"
	"path/filepath"
)

var runEnv *string

func cmdRun(dir string) *cli.Command {
	return &cli.Command{
		Name:  constants2.RunName,
		Usage: constants2.RunUsage,
		Flags: []cli.Flag{
			flags.NewFlag(flags.Env, flags.WithDestination(runEnv), flags.Required()),
		},
		Action: func(ctx *cli.Context) error {
			env := ctx.String(constants2.EnvName)

			cfgLoc := filepath.Join(dir, "env", env, "config.yaml")

			return exec.Command("wyvrn", fmt.Sprintf("-cfg=%s", cfgLoc)).Run()
		},
	}
}
