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
