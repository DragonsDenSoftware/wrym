package main

import (
	"github.com/DragonsDenSoftware/wrym/internal/app/commands"
	"github.com/syke99/trier"
	"log"
	"os"
	"path/filepath"

	"github.com/urfave/cli/v2"
)

func main() {
	ex, _ := os.Executable()
	curDir := filepath.Dir(ex)

	app := &cli.App{}

	app.Commands = []*cli.Command{
		commands.NewCommand(commands.New, curDir),
		commands.NewCommand(commands.Run, curDir),
	}

	try := trier.NewTrier()

	if try.Try(func(args ...any) error {
		return app.Run(os.Args)
	}).Err() != nil {
		log.Fatal(try.Err())
	}
}
