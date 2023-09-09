package flags

import (
	"github.com/urfave/cli/v2"
)

type Flags int64

const (
	Help Flags = iota
	Module
	Language
	Directory
	Dev
	Staging
	Prod
)

type FlagOptions func(flag cli.Flag)

func WithDestination[D *string | *bool](destination D) func(cli.Flag) {
	return func(flag cli.Flag) {
		switch any(destination).(type) {
		case *string:
			f := flag.(*cli.StringFlag)

			f.Destination = any(destination).(*string)
		case *bool:
			f := flag.(*cli.BoolFlag)

			f.Destination = any(destination).(*bool)
		}
	}
}

func Required() func(cli.Flag) {

	return func(flag cli.Flag) {
		f := flag.(*cli.StringFlag)

		f.Required = true
	}
}

func NewFlag(flag Flags, opts ...FlagOptions) cli.Flag {
	var f cli.Flag
	switch flag {
	case Help:
		f = help()
	case Module:
		f = module()
	case Language:
		f = language()
	case Directory:
		f = directory()
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
