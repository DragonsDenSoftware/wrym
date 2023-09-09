package commands

import (
	"errors"
	"github.com/syke99/wyvrn-cli/internal/constants"
	"github.com/syke99/wyvrn-cli/internal/flags"
	"github.com/urfave/cli/v2"
	"os"
	"strings"
)

var newHelp *bool
var module *string
var language *string
var directory *string

func New() *cli.Command {
	return &cli.Command{
		Name:  constants.NewName,
		Usage: constants.NewUsage,
		Flags: []cli.Flag{
			flags.NewFlag(flags.Help, flags.WithDestination(newHelp)),
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
	var err error

	if _, ok := constants.Languages[strings.ToLower(*language)]; !ok {
		err = errors.New("language entered not in list of supported languages")
	}

	if err == nil {
		err = os.Chdir(*directory)
	}

	return nil
}
