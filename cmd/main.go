package main

import (
	"github.com/DragonsDenSoftware/wrym/internal/app/commands"
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

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
