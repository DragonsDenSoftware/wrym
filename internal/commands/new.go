package commands

import (
	"github.com/syke99/wyvrn-cli/internal/flags"
	"github.com/urfave/cli/v2"
)

var module *string
var language *string
var directory *string

func New() *cli.Command {
	return &cli.Command{
		Name:  "new",
		Usage: "generate a new wyvrn app/module",
		Flags: []cli.Flag{
			flags.NewFlag(flags.Module, flags.WithDestination(module)),
			flags.NewFlag(flags.Language, flags.WithDestination(language), flags.Required()),
			flags.NewFlag(flags.Directory, flags.WithDestination(directory)),
		},
		Action: func(c *cli.Context) error {
			return create()
		},
	}
}

func create() error {
	// TODO: fill body with creating new apps/modules
	return nil
}
