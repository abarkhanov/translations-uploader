package app

import (
	"os"

	"github.com/abarkhanov/ttu/internal/client"
	"github.com/abarkhanov/ttu/internal/config"
	"github.com/abarkhanov/ttu/internal/uploader"
)

func uploadTranslations() error {
	cfg, err := config.Load()
	if err != nil {
		return err
	}

	err = validateConfig(cfg)
	if err != nil {
		return err
	}

	isReadyForStart, err := confirmConfig(cfg)
	if err != nil {
		return err
	}
	if isReadyForStart == false {
		os.Exit(0)
	}

	apiClient := client.Init(cfg.TargetAPIAuthorizationKey, cfg.TargetAPIHost)
	err = uploader.LoadTranslations(apiClient, cfg)
	if err != nil {
		return err
	}

	return nil
}
