package flags

import (
	"fmt"
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
	Config
	Env
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
						return fmt.Errorf("%s flag is required", s)
					}

					if flg == constants.NewName {
						return fmt.Errorf("%s flag required when using %s command", s, constants.NewName)
					}

					if flg == constants.RunName {
						return fmt.Errorf("%s flag required when using %s command", s, constants.RunName)
					}

					return fmt.Errorf("%s flag required when using %s flag", s, flg)
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
		case constants.ModuleName:
			if ctx.String(flg) != "" {
				flag = constants.ModuleName
				required = true
			}
		case constants.NewName:
			if ctx.String(constants.ModuleName) == "" &&
				ctx.Command.Name == constants.NewName {
				flag = constants.NewName
				required = true
			}
		case constants.RunName:
			if ctx.Command.Name == constants.RunName {
				flag = constants.RunName
				required = true
			}
		case constants.ConfigName:
			if ctx.Bool(constants.ConfigName) {
				flag = constants.ConfigName
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
	}

	for _, opt := range opts {
		opt(f)
	}

	return f
}
