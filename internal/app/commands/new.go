package commands

import (
	"errors"
	"fmt"
	"github.com/DragonsDenSoftware/wrym/internal/app/flags"
	"github.com/DragonsDenSoftware/wrym/internal/pkg/constants"
	"github.com/DragonsDenSoftware/wrym/internal/pkg/templates"
	"github.com/iancoleman/strcase"
	"github.com/syke99/trier"
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
var app *bool

func cmdNew(homeDir string) *cli.Command {
	return &cli.Command{
		Name:  constants.NewName,
		Usage: constants.NewUsage,
		Flags: []cli.Flag{
			flags.NewFlag(flags.App, flags.WithDestination(app)),
			flags.NewFlag(flags.Module, flags.WithDestination(module)),
			flags.NewFlag(flags.Language, flags.WithDestination(language), flags.Required()),
			flags.NewFlag(flags.Directory, flags.WithDestination(directory)),
			flags.NewFlag(flags.Step, flags.WithDestination(step), flags.Required(constants.ModuleName, constants.AppName)),
			flags.NewFlag(flags.Name, flags.WithDestination(name), flags.Required(constants.AppName)),
			flags.NewFlag(flags.Config, flags.WithDestination(cfg)),
			flags.NewFlag(flags.Env, flags.WithDestination(newEnv), flags.Required(constants.ConfigName)),
		},
		Action: func(c *cli.Context) error {
			return create(homeDir)
		},
	}
}

func create(homeDir string) error {
	var err error

	lang := strings.ToLower(*language)

	tr := trier.NewTrier()

	tr.Try(func(args ...any) error {
		if _, ok := constants.Languages[lang]; !ok {
			err = errors.New("language entered not in list of supported languages")
		}
		return err
	}).Try(func(args ...any) error {
		try := trier.NewTrier()

		// if specific directory was specified
		if directory != nil {
			try.Try(func(args ...any) error {
				homeDir, err = changeToDir(homeDir)
				return err
			})
		} else if app != nil {
			try.Try(func(args ...any) error {
				return os.Mkdir(*name, os.ModeDir)
			}).Try(func(args ...any) error {
				return os.Chdir(*name)
			}).Try(func(args ...any) error {
				homeDir = filepath.Join(homeDir, *name)
				return nil
			})
		}

		return try.Err()
	}).Try(func(args ...any) error {
		try := trier.NewTrier()

		if module != nil {
			try.Try(func(args ...any) error {
				return templateModule(homeDir, lang)
			})
		} else if cfg != nil {
			var envFile *os.File
			var e error
			try.Try(func(args ...any) error {
				_, e = os.Stat(filepath.Join(homeDir, "env"))
				return e
			}).TryJoin(func(args ...any) error {
				_, e = os.Stat(filepath.Join(homeDir, "env", *newEnv))
				return e
			}).Try(func(args ...any) error {
				envFile, e = os.Create(filepath.Join(homeDir, "env", *newEnv))
				return e
			}).Try(func(args ...any) error {
				_, e = envFile.WriteString(templates.Config)
				_ = envFile.Close()
				return e
			})
		} else {
			try.Try(func(args ...any) error {
				return os.Mkdir("modules", os.ModeDir)
			}).Try(func(args ...any) error {
				return os.Mkdir("env", os.ModeDir)
			}).Try(func(args ...any) error {
				return os.Mkdir(filepath.Join("env", constants.DevEnv), os.ModeDir)
			}).Try(func(args ...any) error {
				return os.Mkdir(filepath.Join("env", constants.StageEnv), os.ModeDir)
			}).Try(func(args ...any) error {
				return os.Mkdir(filepath.Join("env", constants.ProdEnv), os.ModeDir)
			}).Try(func(args ...any) error {
				return templateModule(filepath.Join(homeDir, "modules"), lang)
			})
		}

		return try.Err()
	})

	return err
}

func changeToDir(homeDir string) (string, error) {
	try := trier.NewTrier()

	cleanDir := filepath.Clean(*directory)

	try.Try(func(args ...any) error {
		_, err := os.Stat(cleanDir)
		return err
	}).Try(func(args ...any) error {
		return os.Chdir(cleanDir)
	}).Try(func(args ...any) error {
		homeDir = cleanDir
		return nil
	})

	return homeDir, try.Err()
}

func templateModule(dir string, lang string) error {
	try := trier.NewTrier()

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

	try.Try(func(args ...any) error {
		var e error
		f, e = os.Create(filepath.Join(dir, mod, ext))
		return e
	}).Try(func(args ...any) error {
		return templates.Template(f, tmplLang, step)
	})

	return try.Err()
}
