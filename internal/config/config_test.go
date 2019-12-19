package config

import (
	"os"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"testing"
)

func TestLoad(t *testing.T) {
	path := "./translations"
	key := "1234qwer"
	host := "http://example.com"

	err := os.Setenv("TRANSLATIONS_PATH", path)
	require.NoError(t, err)
	err = os.Setenv("TARGET_API_AUTHORIZATION_KEY", key)
	require.NoError(t, err)
	err = os.Setenv("TARGET_API_HOST", host)
	require.NoError(t, err)

	c, err := Load()
	require.NoError(t, err)

	assert.Equal(t, path, c.TranslationsPath)
	assert.Equal(t, key, c.TargetAPIAuthorizationKey)
	assert.Equal(t, host, c.TargetAPIHost)
}

func TestLoadNoAPIHost(t *testing.T) {
	path := "./translations"
	key := "1234qwer"

	err := os.Setenv("TRANSLATIONS_PATH", path)
	require.NoError(t, err)
	err = os.Setenv("TARGET_API_AUTHORIZATION_KEY", key)
	require.NoError(t, err)

	_, err = Load()
	require.Error(t, err)
	assert.Equal(t, "required key TARGET_API_HOST missing value", err.Error())
}

func TestLoadNoAPIKey(t *testing.T) {
	path := "./translations"
	host := "http://example.com"

	err := os.Setenv("TRANSLATIONS_PATH", path)
	require.NoError(t, err)
	err = os.Setenv("TARGET_API_HOST", host)
	require.NoError(t, err)

	_, err = Load()
	require.Error(t, err)
	assert.Equal(t, "required key TARGET_API_AUTHORIZATION_KEY missing value", err.Error())
}
