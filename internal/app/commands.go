package app

import (
	"github.com/abarkhanov/ttu/internal/client"
	"github.com/abarkhanov/ttu/internal/config"
	"github.com/abarkhanov/ttu/internal/uploader"
)

func uploadTranslations(cfg *config.Config) error {
	apiClient := client.Init(cfg.TargetAPIKey, cfg.TargetAPIHost)
	err := uploader.LoadTranslations(apiClient, cfg)
	if err != nil {
		return err
	}

	return nil
}
