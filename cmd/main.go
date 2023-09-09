package cmd

import (
	"log"
	"os"

	"github.com/syke99/wyvrn-cli/internal/commands"
	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Commands: []*cli.Command{
			commands.New(),
			commands.Run(),
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
