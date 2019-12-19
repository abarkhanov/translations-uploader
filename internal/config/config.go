package config

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/kelseyhightower/envconfig"
)

// Config contains service configuration settings
type Config struct {
	TranslationsPath          string `envconfig:"TRANSLATIONS_PATH"`
	TargetAPIAuthorizationKey string `envconfig:"TARGET_API_AUTHORIZATION_KEY"`
	TargetAPIHost             string `envconfig:"TARGET_API_HOST"`
	OrgIDSNCF                 string `envconfig:"ORGID_SNCF"`
	OrgIDThalys               string `envconfig:"ORGID_THALYS"`
}

// Load parses environment variables returning configuration
func Load() (*Config, error) {
	cfg := &Config{}
	err := envconfig.Process("", cfg)
	if err != nil {
		return nil, err
	}

	if cfg.TranslationsPath == "" {
		v, err := promtParameter("TRANSLATIONS_PATH", true)
		if err != nil {
			return nil, err
		}

		cfg.TranslationsPath = v
	}

	if cfg.TargetAPIAuthorizationKey == "" {
		v, err := promtParameter("TARGET_API_AUTHORIZATION_KEY", true)
		if err != nil {
			return nil, err
		}
		cfg.TargetAPIAuthorizationKey = v
	}

	if cfg.TargetAPIHost == "" {
		v, err := promtParameter("TARGET_API_HOST", true)
		if err != nil {
			return nil, err
		}
		cfg.TargetAPIHost = v
	}

	if cfg.OrgIDSNCF == "" {
		v, err := promtParameter("ORGID_SNCF", false)
		if err != nil {
			return nil, err
		}
		if v == "" {
			fmt.Println(fmt.Sprintf("Note! Translations won't be uploaded for SNCF"))
		}

		cfg.OrgIDSNCF = v
	}

	if cfg.OrgIDThalys == "" {
		v, err := promtParameter("ORGID_THALYS", false)
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

func Confirm(cfg *Config) (isReadyForStart bool, err error) {
	isReadyForStart = false
	fmt.Println("================================================================================================")
	fmt.Println("Check your config: ")
	fmt.Println("================================================================================================")
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

func Validate(cfg *Config) error {
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

func promtParameter(paramName string, required bool) (string, error) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(fmt.Sprintf("Enter %s: ", paramName))
	v, err := reader.ReadString('\n')
	v = strings.TrimSpace(strings.TrimSuffix(v, "\n"))
	if err != nil {
		return "", err
	}

	if required == true && len(v) == 0 {
		return "", errors.New(fmt.Sprintf("%s is required parameter", paramName))
	}

	if required == false && len(v) == 0 {
		fmt.Println(fmt.Sprintf("%s is not set", paramName))
		return "", nil
	}

	return v, nil
}
