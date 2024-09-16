package uploader

import (
	"github.com/abarkhanov/ttu/internal/config/upload_config"
	"os"
	"testing"

	"github.com/abarkhanov/ttu/testing/mocks"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestLoadTranslations(t *testing.T) {
	path := "./translations"
	key := "1234qwer"
	host := "http://example.com"
	sncfID := "sncf-1234"
	thalysID := "thalys-9999"

	err := os.Setenv("TRANSLATIONS_PATH", path)
	require.NoError(t, err)
	err = os.Setenv("TARGET_API_AUTHORIZATION_KEY", key)
	require.NoError(t, err)
	err = os.Setenv("TARGET_API_HOST", host)
	require.NoError(t, err)
	err = os.Setenv("ORGID_SNCF", sncfID)
	require.NoError(t, err)
	err = os.Setenv("ORGID_THALYS", thalysID)
	require.NoError(t, err)

	c, err := upload_config.LoadUploadCfg()
	require.NoError(t, err)

	client := &mocks.ApiTranslationsClient{}
	client.Mock.On("AddToken", mock.AnythingOfType("string"), mock.AnythingOfType("string"), mock.AnythingOfType("string"), mock.Anything).
		Return(nil)
	err = LoadTranslations(client, c)
	require.NoError(t, err)
}
