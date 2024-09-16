package download_config

import (
	"bufio"
	"errors"
	"fmt"
	"github.com/abarkhanov/ttu/internal/config"
	"github.com/kelseyhightower/envconfig"
	"os"
	"strings"
)

// UploadConfig contains service configuration settings
type DownloadConfig struct {
	SourceAPIAuthorizationKey string `envconfig:"SOURCE_API_AUTHORIZATION_KEY" required:"true"`
	SourceAPIHost             string `envconfig:"SOURCE_API_HOST" required:"true"`
	OrgID                     string `envconfig:"ORG_ID"`
	OrgName                   string `envconfig:"ORG_NAME"`
}

func Load() (*DownloadConfig, error) {
	cfg := &DownloadConfig{}
	err := envconfig.Process("", cfg)
	if err != nil {
		return nil, err
	}

	if cfg.SourceAPIAuthorizationKey == "" {
		v, err := config.PromtParameter("SOURCE_API_AUTHORIZATION_KEY", true)
		if err != nil {
			return nil, err
		}
		cfg.SourceAPIAuthorizationKey = v
	}

	if cfg.SourceAPIHost == "" {
		v, err := config.PromtParameter("SOURCE_API_HOST", true)
		if err != nil {
			return nil, err
		}
		cfg.SourceAPIHost = v
	}

	if cfg.OrgID == "" {
		v, err := config.PromtParameter("ORG_ID", false)
		if err != nil {
			return nil, err
		}
		if v != "" {
			cfg.OrgID = v
		}
	}

	if cfg.OrgName == "" {
		v, err := config.PromtParameter("ORG_NAME", false)
		if err != nil {
			return nil, err
		}
		if v != "" {
			cfg.OrgName = v
		}
	}

	return cfg, nil
}

func Validate(cfg *DownloadConfig) error {
	if cfg.SourceAPIAuthorizationKey == "" {
		return errors.New("parameter SOURCE_API_AUTHORIZATION_KEY not set")

	}

	if cfg.SourceAPIHost == "" {
		return errors.New("parameter SOURCE_API_HOST not set")
	}

	if (cfg.OrgName == "" && cfg.OrgID != "") || (cfg.OrgName != "" && cfg.OrgID == "") {
		return errors.New("parameters ORG_ID and ORG_NAME should be both set or both empty")
	}

	return nil
}

func Confirm(cfg *DownloadConfig) (isReadyForStart bool, err error) {
	isReadyForStart = false
	fmt.Println("================================================================================================")
	fmt.Println("Check your config: ")
	fmt.Println("================================================================================================")
	fmt.Println(fmt.Sprintf("SOURCE_API_HOST              : %s ", cfg.SourceAPIHost))
	fmt.Println(fmt.Sprintf("SOURCE_API_AUTHORIZATION_KEY : %s ", cfg.SourceAPIAuthorizationKey))
	fmt.Println(fmt.Sprintf("ORG_ID                       : %s ", cfg.OrgID))
	fmt.Println(fmt.Sprintf("ORG_NAME                     : %s ", cfg.OrgName))
	fmt.Println("================================================================================================")
	fmt.Println("Translations will be downloaded in ./translations dir")
	fmt.Println("================================================================================================")
	if cfg.OrgID == "" && cfg.OrgName == "" {
		fmt.Println("Note! Only default translations Will be downloaded")
	}
	if cfg.OrgID != "" && cfg.OrgName != "" {
		fmt.Println("Note! Will be downloaded default translations and translations for Org: " + cfg.OrgName)
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
