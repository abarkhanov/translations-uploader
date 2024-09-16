package download_config

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"os"
	"testing"
)

func TestLoadDownloadDefaultCfg(t *testing.T) {
	key := "1234qwer"
	host := "http://example.com"

	err := os.Setenv("SOURCE_API_AUTHORIZATION_KEY", key)
	require.NoError(t, err)
	err = os.Setenv("SOURCE_API_HOST", host)
	require.NoError(t, err)
	err = os.Setenv("ORG_ID", "'\n'")
	require.NoError(t, err)
	err = os.Setenv("ORG_NAME", "'\n'")
	require.NoError(t, err)

	c, err := Load()
	require.NoError(t, err)

	assert.Equal(t, key, c.SourceAPIAuthorizationKey)
	assert.Equal(t, host, c.SourceAPIHost)
}

func TestLoadDownloadCfgWithOrg(t *testing.T) {
	key := "1234qwer"
	host := "http://example.com"
	orgID := "123-456-qwer-456"
	orgName := "sncf"

	err := os.Setenv("SOURCE_API_AUTHORIZATION_KEY", key)
	require.NoError(t, err)
	err = os.Setenv("SOURCE_API_HOST", host)
	require.NoError(t, err)
	err = os.Setenv("ORG_ID", orgID)
	require.NoError(t, err)
	err = os.Setenv("ORG_NAME", orgName)
	require.NoError(t, err)

	c, err := Load()
	require.NoError(t, err)

	assert.Equal(t, key, c.SourceAPIAuthorizationKey)
	assert.Equal(t, host, c.SourceAPIHost)
	assert.Equal(t, orgID, c.OrgID)
	assert.Equal(t, orgName, c.OrgName)
}

func TestLoadDownloadCfgWithError(t *testing.T) {
	_, err := Load()
	require.Error(t, err)
}
