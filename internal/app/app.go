package app

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/abarkhanov/ttu/internal/config"
	"github.com/urfave/cli/v2"
)

func New() (*cli.App, error) {
	app := &cli.App{
		Name:     "Tpl email translations uploader",
		Usage:    "To start uploading run: $> ttu upload",
		Commands: initCommands(),
	}
	app.Name = "Translation uploader"
	app.Version = "0.0.1"
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

func confirmConfig(cfg *config.Config) (isReadyForStart bool, err error) {
	isReadyForStart = false
	fmt.Println("================================================================================================")
	fmt.Println("Check your config: ")
	fmt.Println(fmt.Sprintf("TRANSLATIONS_PATH            : %s ", cfg.TranslationsPath))
	fmt.Println(fmt.Sprintf("TARGET_API_HOST              : %s ", cfg.TargetAPIHost))
	fmt.Println(fmt.Sprintf("ORGID_SNCF                   : %s ", cfg.OrgIDSNCF))
	fmt.Println(fmt.Sprintf("ORGID_THALYS                 : %s ", cfg.OrgIDThalys))
	fmt.Println(fmt.Sprintf("TARGET_API_AUTHORIZATION_KEY : %s ", cfg.TargetAPIAuthorizationKey))
	fmt.Println("================================================================================================")
	if cfg.OrgIDSNCF == "" {
		fmt.Println("Note! Translations won't be uploaded for SNCF")
	}
	if cfg.OrgIDThalys == "" {
		fmt.Println("Note! Translations won't be uploaded for Thalys")
	}
	if cfg.OrgIDSNCF == "" && cfg.OrgIDThalys == "" {
		fmt.Println("Note! Will be uploaded only default translations")
	}
	fmt.Println("================================================================================================")
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Do you want to start? (y/n):")
	v, err := reader.ReadString('\n')
	if err != nil {
		return isReadyForStart, err
	}
	v = strings.TrimSpace(strings.TrimSuffix(v, "\n"))

	if v == "y" {
		isReadyForStart = true
	}

	if v == "n" {
		isReadyForStart = false
	}

	if isReadyForStart == false {
		fmt.Println("Exit...")
	} else {
		fmt.Println("Starting...")
	}

	return isReadyForStart, nil
}

func validateConfig(cfg *config.Config) error {
	if cfg.TranslationsPath == "" {
		return errors.New("parameter TRANSLATIONS_PATH not set")
	}

	if cfg.TargetAPIAuthorizationKey == "" {
		return errors.New("parameter TARGET_API_AUTHORIZATION_KEY not set")

	}

	if cfg.TargetAPIHost == "" {
		return errors.New("parameter TARGET_API_HOST not set")
	}

	return nil
}
