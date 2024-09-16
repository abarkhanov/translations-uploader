package upload_config

import (
	"bufio"
	"errors"
	"fmt"
	"github.com/abarkhanov/ttu/internal/config"
	"os"
	"strings"

	"github.com/kelseyhightower/envconfig"
)

// UploadConfig contains service configuration settings
type UploadConfig struct {
	TranslationsPath          string `envconfig:"SOURCE_TRANSLATIONS_PATH" required:"true"`
	TargetAPIAuthorizationKey string `envconfig:"TARGET_API_AUTHORIZATION_KEY" required:"true"`
	TargetAPIHost             string `envconfig:"TARGET_API_HOST" required:"true"`
	OrgIDSNCF                 string `envconfig:"ORGID_SNCF"`
	OrgIDThalys               string `envconfig:"ORGID_THALYS"`
}

// LoadUploadCfg parses environment variables returning configuration
func LoadUploadCfg() (*UploadConfig, error) {
	cfg := &UploadConfig{}
	err := envconfig.Process("", cfg)
	if err != nil {
		return nil, err
	}

	if cfg.OrgIDSNCF == "" {
		v, err := config.PromtParameter("ORGID_SNCF", false)
		if err != nil {
			return nil, err
		}
		if v == "" {
			fmt.Println(fmt.Sprintf("Note! Translations won't be uploaded for SNCF"))
		}

		cfg.OrgIDSNCF = v
	}

	if cfg.OrgIDThalys == "" {
		v, err := config.PromtParameter("ORGID_THALYS", false)
		if err != nil {
			return nil, err
		}
		if v == "" {
			fmt.Println(fmt.Sprintf("Note! Translations won't be uploaded for THALYS"))
		}

		cfg.OrgIDThalys = v
	}

	return cfg, nil
}

func Confirm(cfg *UploadConfig) (isReadyForStart bool, err error) {
	isReadyForStart = false
	fmt.Println("================================================================================================")
	fmt.Println("Check your config: ")
	fmt.Println("================================================================================================")
	fmt.Println(fmt.Sprintf("SOURCE_TRANSLATIONS_PATH            : %s ", cfg.TranslationsPath))
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

func Validate(cfg *UploadConfig) error {
	if cfg.TranslationsPath == "" {
		return errors.New("parameter SOURCE_TRANSLATIONS_PATH not set")
	}

	if cfg.TargetAPIAuthorizationKey == "" {
		return errors.New("parameter TARGET_API_AUTHORIZATION_KEY not set")

	}

	if cfg.TargetAPIHost == "" {
		return errors.New("parameter TARGET_API_HOST not set")
	}

	return nil
}
