package flags

import "github.com/urfave/cli/v2"

type Flags int64

const (
	Module Flags = iota
	Language
	Directory
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

func Required() func(*cli.StringFlag) {
	return func(flag *cli.StringFlag) {
		flag.Required = true
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
