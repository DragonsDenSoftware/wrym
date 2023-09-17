package commands

import (
	"github.com/urfave/cli/v2"
)

type Commands int8

const (
	New Commands = iota
	Run
)

func NewCommand(cmd Commands, homeDir string) *cli.Command {
	var c *cli.Command
	switch cmd {
	case New:
		c = cmdNew(homeDir)
	case Run:
		c = cmdRun(homeDir)
	}
	return c
}
