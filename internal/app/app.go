package app

import (
	"fmt"
	"log"
	"time"

	"github.com/urfave/cli/v2"
)

func New() (*cli.App, error) {
	app := &cli.App{
		Commands: initCommands(),
	}
	app.Name = "Template Translations Uploader"
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
			Usage:   "Upload translations from files to destination API: $>ttu upload",
			Action: func(c *cli.Context) error {
				err := uploadTranslations()
				if err != nil {
					log.Fatalf("Unable to execute command %s", err)
				}
				return nil
			},
		},
		{
			Name:    "download",
			Aliases: []string{"d"},
			Usage:   "Download translations from origin API to yaml files: $>ttu download",
			Action: func(c *cli.Context) error {
				err := downloadTranslations()
				if err != nil {
					fmt.Println("Unable to execute command", err)
					return err
				}
				return nil
			},
		},
	}

	return commands
}
