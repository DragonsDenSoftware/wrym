package flags

import "github.com/urfave/cli/v2"

func language() *cli.StringFlag {
	return &cli.StringFlag{
		Name:    "language",
		Aliases: []string{"l"},
		Value:   "",
		Usage:   "the language you want to create your new wyvrn app/module for; apps can run modules compiled from multiple languages, but needs a language to make the first module with",
	}
}
