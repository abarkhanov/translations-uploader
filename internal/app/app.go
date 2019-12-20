package app

import (
	"log"
	"time"

	"github.com/urfave/cli/v2"
)

func New() (*cli.App, error) {
	app := &cli.App{
		Commands: initCommands(),
	}
	app.Name = "Template Translations Uploader"
	app.Usage = "To start uploading run: $> ttu upload"
	app.Version = "0.0.2 (Yury's build)"
	app.Compiled = time.Now()
	app.EnableBashCompletion = true
	app.HideHelp = false
	app.HideVersion = false

	return app, nil
}

func initCommands() []*cli.Command {
	commands := []*cli.Command{
		{
			Name:    "upload",
			Aliases: []string{"u"},
			Usage:   "Upload translations from files to destination API",
			Action: func(c *cli.Context) error {
				err := uploadTranslations()
				if err != nil {
					log.Fatalf("Unable to execute command %s", err)
				}
				return nil
			},
		},
	}

	return commands
}
