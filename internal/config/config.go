package config

import (
	"github.com/kelseyhightower/envconfig"
)

// Config contains service configuration settings
type Config struct {
	TranslationsPath string `envconfig:"TRANSLATIONS_PATH" required:"true"`
	TargetAPIKey     string `envconfig:"TARGET_API_KEY" required:"true"`
	TargetAPIHost    string `envconfig:"TARGET_API_HOST" required:"true"`
	OrgIDSNCF        string `envconfig:"ORGID_SNCF"`
	OrgIDThalys      string `envconfig:"ORGID_THALYS"`
}

// Load parses environment variables returning configuration
func Load() (*Config, error) {
	cfg := &Config{}
	err := envconfig.Process("", cfg)
	if err != nil {
		return nil, err
	}

	return cfg, nil
}
