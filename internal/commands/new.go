package commands

import (
	"errors"
	"fmt"
	"github.com/iancoleman/strcase"
	"github.com/syke99/wyvrn-cli/internal/constants"
	"github.com/syke99/wyvrn-cli/internal/flags"
	"github.com/syke99/wyvrn-cli/internal/templates"
	"github.com/urfave/cli/v2"
	"os"
	"strings"
)

var module *string
var language *string
var directory *string
var step *string

func cmdNew(ctx *cli.Context) *cli.Command {
	return &cli.Command{
		Name:  constants.NewName,
		Usage: constants.NewUsage,
		Flags: []cli.Flag{
			flags.NewFlag(flags.Module, flags.WithDestination(module)),
			flags.NewFlag(flags.Language, flags.WithDestination(language), flags.Required(nil)),
			flags.NewFlag(flags.Directory, flags.WithDestination(directory)),
			flags.NewFlag(flags.Step, flags.WithDestination(step), flags.Required(ctx)),
		},
		Action: func(c *cli.Context) error {
			return create()
		},
	}
}

func create() error {
	// TODO: fill body with creating new apps/modules
	var err error

	lang := strings.ToLower(*language)

	if _, ok := constants.Languages[lang]; !ok {
		err = errors.New("language entered not in list of supported languages")
	}

	if err == nil {
		err = os.Chdir(*directory)
	}

	if module != nil {
		var tmplLang templates.Language
		var f *os.File

		mod := *module

		ext := ""

		switch lang {
		case "c":
			tmplLang = templates.C

			mod = strcase.ToKebab(mod)
			ext = "c"
		case "haskell":
			tmplLang = templates.Haskell

			mod = strcase.ToCamel(mod)
			ext = "hs"
		case "rust":
			tmplLang = templates.Rust

			mod = strings.ToLower(mod)
			ext = "rs"
		//case "go":
		//	tmplLang = templates.Go

		//	mod = strcase.ToSnake(mod)
		//	ext = "go"
		case "assemblyscript":
		case "assembly-script":
			tmplLang = templates.AssemblyScript

			mod = strings.ToLower(mod)
			ext = "ts"
		//case "csharp":
		//case "c-sharp":
		//	tmplLang = templates.CSharp

		case "zig":
			tmplLang = templates.Zig

			mod = strings.ToLower(mod)
			ext = "zig"
		case "javascript":
		case "js":
			tmplLang = templates.JavaScript

			mod = strings.ToLower(fmt.Sprintf("%s.%s", mod, ext))
			ext = "js"
			//case "python":
			//tmplLang = templates.Python

		}

		f, err = os.Create(*module)

		if err != nil {
			return err
		}

		err = templates.Template(f, tmplLang, step)
	}

	return err
}
