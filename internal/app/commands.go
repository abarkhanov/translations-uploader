package app

import (
	"github.com/abarkhanov/ttu/internal/config/download_config"
	"github.com/abarkhanov/ttu/internal/config/upload_config"
	"github.com/abarkhanov/ttu/internal/downloader"
	"os"

	"github.com/abarkhanov/ttu/internal/client"
	"github.com/abarkhanov/ttu/internal/uploader"
)

func uploadTranslations() error {
	cfg, err := upload_config.LoadUploadCfg()
	if err != nil {
		return err
	}

	err = upload_config.Validate(cfg)
	if err != nil {
		return err
	}

	isReadyForStart, err := upload_config.Confirm(cfg)
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

func downloadTranslations() error {
	cfg, err := download_config.Load()
	if err != nil {
		return err
	}

	err = download_config.Validate(cfg)
	if err != nil {
		return err
	}

	isReadyForStart, err := download_config.Confirm(cfg)
	if err != nil {
		return err
	}
	if isReadyForStart == false {
		os.Exit(0)
	}

	apiClient := client.Init(cfg.SourceAPIAuthorizationKey, cfg.SourceAPIHost)
	err = downloader.LoadTranslations(apiClient, cfg)
	if err != nil {
		return err
	}

	return nil
}
