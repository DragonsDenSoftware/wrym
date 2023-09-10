package main

import (
	"context"
	"github.com/syke99/wyvrn-cli/internal/commands"
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

var help *bool

func main() {
	ctx := context.Background()

	app := &cli.App{}

	cCtx := cli.NewContext(app, nil, &cli.Context{
		Context: ctx,
		App:     app,
	})

	app.Commands = []*cli.Command{
		commands.NewCommand(cCtx, commands.New),
		commands.NewCommand(cCtx, commands.Run),
	}

	err := app.RunContext(ctx, os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
