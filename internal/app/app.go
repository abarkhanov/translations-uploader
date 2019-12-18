package app

import (
	"fmt"
	"log"

	"github.com/abarkhanov/ttu/internal/config"
	"github.com/urfave/cli/v2"
)

func New() *cli.App {
	cfg, err := config.Load()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("123")

	app := &cli.App{
		Name:     "Tpl email translations uploader",
		Usage:    "fight the loneliness!",
		Commands: initCommands(cfg),
	}

	return app

}

func initCommands(cfg *config.Config) []*cli.Command {
	commands := []*cli.Command{
		{
			Name:    "upload",
			Aliases: []string{"u"},
			Usage:   "Upload translations from files to destination API of translations service",
			Action: func(c *cli.Context) error {
				err := uploadTranslations(cfg)
				if err != nil {
					log.Fatalf("Unable to execute command %s", err)
				}
				return nil
			},
		},
	}

	return commands
}
