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
	Name
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

func Required(ctx *cli.Context, reqFor ...string) func(*cli.StringFlag) {
	return func(flag *cli.StringFlag) {
		if (len(reqFor) == 1 && reqFor[0] == "all") ||
			isRequired(ctx, reqFor) {
			flag.Required = true
			return
		}
	}
}

func isRequired(ctx *cli.Context, reqFor []string) bool {
	var required bool

	for _, flg := range reqFor {
		switch flg {
		case constants.ModuleName:
			if ctx.String(flg) != "" {
				required = true
			}
		case constants.NewName:
			if ctx.String(constants.ModuleName) == "" {
				required = true
			}
		}
	}

	return required
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
	case Name:
		f = name()
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
