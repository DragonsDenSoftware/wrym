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
var cfg *bool
var newEnv *string

func cmdNew(homeDir string) *cli.Command {
	return &cli.Command{
		Name:  constants.NewName,
		Usage: constants.NewUsage,
		Flags: []cli.Flag{
			flags.NewFlag(flags.Module, flags.WithDestination(module)),
			flags.NewFlag(flags.Language, flags.WithDestination(language), flags.Required()),
			flags.NewFlag(flags.Directory, flags.WithDestination(directory)),
			flags.NewFlag(flags.Step, flags.WithDestination(step), flags.Required(constants.ModuleName)),
			flags.NewFlag(flags.AppName, flags.WithDestination(name), flags.Required(constants.NewName)),
			flags.NewFlag(flags.Config, flags.WithDestination(cfg)),
			flags.NewFlag(flags.Env, flags.WithDestination(newEnv), flags.Required(constants.ConfigName)),
		},
		Action: func(c *cli.Context) error {
			return create(homeDir)
		},
	}
}

func create(homeDir string) error {
	// TODO: rework to consider option of creating a new env
	var err error

	lang := strings.ToLower(*language)

	if _, ok := constants.Languages[lang]; !ok {
		err = errors.New("language entered not in list of supported languages")
	}

	if err == nil {
		// if specific directory was specified
		if directory != nil {
			homeDir, err = changeToDir(homeDir)
		} else if name != nil {
			err = os.Mkdir(*name, os.ModeDir)

			if err == nil {
				err = os.Chdir(*name)
			}

			if err == nil {
				homeDir = filepath.Join(homeDir, *name)
			}
		}
	}

	if module != nil {
		err = templateModule(homeDir, lang)
	} else if cfg != nil {
		_, err = os.Stat(filepath.Join(homeDir, "env"))

		var envFile *os.File
		if err == nil {
			if _, err = os.Stat(filepath.Join(homeDir, "env", *newEnv)); err != nil {
				envFile, err = os.Create(filepath.Join(homeDir, "env", *newEnv))
				defer envFile.Close()
			}
		}

		if err == nil {
			_, err = envFile.WriteString(templates.Config)
		}
	} else {
		// create directory structure, then template initial module
		err = os.Mkdir("modules", os.ModeDir)

		if err == nil {
			err = os.Mkdir("env", os.ModeDir)
		}

		if err == nil {
			err = os.Mkdir(filepath.Join("env", constants.DevEnv), os.ModeDir)
		}

		if err == nil {
			err = os.Mkdir(filepath.Join("env", constants.StageEnv), os.ModeDir)
		}

		if err == nil {
			err = os.Mkdir(filepath.Join("env", constants.ProdEnv), os.ModeDir)
		}

		if err == nil {
			err = templateModule(filepath.Join(homeDir, "modules"), lang)
		}
	}

	return err
}

func changeToDir(homeDir string) (string, error) {
	var err error

	cleanDir := filepath.Clean(*directory)

	// check for existence of specified
	// directory
	_, err = os.Stat(cleanDir)

	if err == nil {
		err = os.Chdir(cleanDir)
	}

	if err == nil {
		homeDir = cleanDir
	}

	return homeDir, err
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
	case "assemblyscript", "assembly-script":
		tmplLang = templates.AssemblyScript

		mod = strings.ToLower(mod)
		ext = "ts"
	//case "csharp", "c-sharp":
	//	tmplLang = templates.CSharp

	case "zig":
		tmplLang = templates.Zig

		mod = strings.ToLower(mod)
		ext = "zig"
	case "javascript", "js":
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
