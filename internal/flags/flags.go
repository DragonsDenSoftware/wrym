package flags

import (
	"github.com/syke99/wyvrn-cli/internal/constants"
	"github.com/urfave/cli/v2"
)

type Flags int8

const (
	Module Flags = iota
	Language
	Directory
	Step
	Dev
	Staging
	Prod
)

type FlagOptions func(flag *cli.StringFlag)

func WithDestination(destination *string) func(*cli.StringFlag) {
	return func(flag *cli.StringFlag) {
		flag.Destination = destination
	}
}

func Required(ctx *cli.Context) func(*cli.StringFlag) {
	return func(flag *cli.StringFlag) {
		if ctx == nil ||
			ctx.Value(constants.ModuleName).(string) != "" {
			flag.Required = true
			return
		}
	}
}

func NewFlag(flag Flags, opts ...FlagOptions) cli.Flag {
	var f *cli.StringFlag
	switch flag {
	case Module:
		f = module()
	case Language:
		f = language()
	case Directory:
		f = directory()
	case Step:
		f = step()
	case Dev:
		f = dev()
	case Staging:
		f = staging()
	case Prod:
		f = prod()
	}

	for _, opt := range opts {
		opt(f)
	}

	return f
}
