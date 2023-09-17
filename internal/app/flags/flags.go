package flags

import (
	constants2 "github.com/syke99/wyvrn-cli/internal/pkg/constants"
	"github.com/syke99/wyvrn-cli/internal/pkg/errs"
	"github.com/urfave/cli/v2"
)

type Flags int8

const (
	Module Flags = iota
	Language
	Directory
	Step
	Name
	Config
	Env
	App
)

type FlagOptions func(flag cli.Flag)

func WithDestination[D *string | *bool](destination D) func(cli.Flag) {
	return func(flag cli.Flag) {
		switch any(flag).(type) {
		case *cli.StringFlag:
			flag.(*cli.StringFlag).Destination = any(destination).(*string)
		case *cli.BoolFlag:
			flag.(*cli.BoolFlag).Destination = any(destination).(*bool)
		}
	}
}

func Required(reqFor ...string) func(cli.Flag) {
	return func(flag cli.Flag) {
		if len(reqFor) == 0 {
			flag.(*cli.StringFlag).Required = true
			return
		}

		switch any(flag).(type) {
		case *cli.StringFlag:
			flag.(*cli.StringFlag).Action = func(ctx *cli.Context, s string) error {
				if flg, ok := isRequired(ctx, reqFor); ok &&
					s == "" {
					if flg == "all" {
						return errs.FlagRequired(s)
					}

					if flg == constants2.NewName {
						return errs.FlagRequiredForCommand(s, constants2.NewName)
					}

					if flg == constants2.RunName {
						return errs.FlagRequiredForCommand(s, constants2.RunName)
					}

					return errs.FlagRequiredForOtherFlag(s, flg)
				}
				return nil
			}
		}
	}
}

func isRequired(ctx *cli.Context, reqFor []string) (string, bool) {
	var flag string
	var required bool

	for _, flg := range reqFor {
		switch flg {
		case constants2.ModuleName:
			if ctx.String(flg) != "" {
				flag = constants2.ModuleName
				required = true
			}
		case constants2.NewName:
			if ctx.String(constants2.ModuleName) == "" &&
				ctx.Command.Name == constants2.NewName {
				flag = constants2.NewName
				required = true
			}
		case constants2.RunName:
			if ctx.Command.Name == constants2.RunName {
				flag = constants2.RunName
				required = true
			}
		case constants2.ConfigName:
			if ctx.Bool(constants2.ConfigName) {
				flag = constants2.ConfigName
				required = true
			}
		}
	}

	return flag, required
}

func NewFlag(flag Flags, opts ...FlagOptions) cli.Flag {
	var f cli.Flag
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
	case Config:
		f = config()
	case Env:
		f = env()
	case App:
		f = app()
	}

	for _, opt := range opts {
		opt(f)
	}

	return f
}
