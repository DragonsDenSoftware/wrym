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
	"path/filepath"
	"strings"
)

var module *string
var language *string
var directory *string
var step *string
var name *string

func cmdNew(ctx *cli.Context) *cli.Command {
	return &cli.Command{
		Name:  constants.NewName,
		Usage: constants.NewUsage,
		Flags: []cli.Flag{
			flags.NewFlag(flags.Module, flags.WithDestination(module)),
			flags.NewFlag(flags.Language, flags.WithDestination(language), flags.Required(nil, "all")),
			flags.NewFlag(flags.Directory, flags.WithDestination(directory)),
			flags.NewFlag(flags.Step, flags.WithDestination(step), flags.Required(ctx, constants.ModuleName)),
			flags.NewFlag(flags.Name, flags.WithDestination(name), flags.Required(ctx, constants.NewName)),
		},
		Action: func(c *cli.Context) error {
			return create()
		},
	}
}

func create() error {
	ex, _ := os.Executable()
	curDir := filepath.Dir(ex)

	var err error

	lang := strings.ToLower(*language)

	if _, ok := constants.Languages[lang]; !ok {
		err = errors.New("language entered not in list of supported languages")
	}

	if err == nil {
		// if specific directory was specified
		if directory != nil {
			var dir os.FileInfo

			cleanDir := filepath.Clean(*directory)

			// check for existence of specified
			// directory
			dir, err = os.Stat(cleanDir)

			if err == nil {
				if dir == nil {
					// if the specified directory
					// doesn't exist, create it
					err = os.Mkdir(cleanDir, os.ModeDir)
				}

				// change to the specified directory
				if err == nil {
					err = os.Chdir(cleanDir)
				}
			}

			if err == nil {
				curDir = cleanDir
			}
		} else if name != nil {
			err = os.Mkdir(*name, os.ModeDir)

			if err == nil {
				err = os.Chdir(*name)
			}

			if err == nil {
				curDir = filepath.Join(curDir, *name)
			}
		}
	}

	if module != nil {
		err = templateModule(curDir, lang)
	} else {
		// create directory structure, then template initial module
		err = os.Mkdir("modules", os.ModeDir)

		if err == nil {
			err = os.Mkdir("env", os.ModeDir)
		}

		if err == nil {
			err = os.Mkdir(filepath.Join("env", constants.DevName), os.ModeDir)
		}

		if err == nil {
			err = os.Mkdir(filepath.Join("env", constants.StagingName), os.ModeDir)
		}

		if err == nil {
			err = os.Mkdir(filepath.Join("env", constants.ProdName), os.ModeDir)
		}

		if err == nil {
			err = templateModule(filepath.Join(curDir, "modules"), lang)
		}
	}

	return err
}

func templateModule(dir string, lang string) error {
	var tmplLang templates.Language
	var f *os.File

	var err error

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

	f, err = os.Create(filepath.Join(dir, mod, ext))

	if err != nil {
		return err
	}

	return templates.Template(f, tmplLang, step)
}
