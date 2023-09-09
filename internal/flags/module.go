package flags

import "github.com/urfave/cli/v2"

func module() *cli.StringFlag {
	return &cli.StringFlag{
		Name:    "module",
		Aliases: []string{"m"},
		Value:   "",
		Usage:   "signal cli to create a template for a WASM module; value will be used to name the module",
	}
}
