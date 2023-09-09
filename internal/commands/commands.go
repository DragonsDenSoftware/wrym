package commands

import "github.com/urfave/cli/v2"

type Commands int8

const (
	New Commands = iota
	Run
)

func NewCommand(cmd Commands) *cli.Command {
	var c *cli.Command
	switch cmd {
	case New:
		c = cmdNew()
	case Run:
		c = cmdRun()
	}
	return c
}
