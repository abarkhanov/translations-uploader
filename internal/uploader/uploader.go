package uploader

import (
	"fmt"
	"io/ioutil"
	"log"
	"path/filepath"
	"strings"

	"github.com/abarkhanov/ttu/internal/config"
	uuid "github.com/satori/go.uuid"
	"gopkg.in/yaml.v2"
)

var blacklistedFiles = map[string]struct{}{
	"cancellation-reasons.yaml": struct{}{},
	"locale-config.yaml":        struct{}{},
	"vehicle-class.yaml":        struct{}{},
	"sms.yaml":                  struct{}{},
}

type translationsList map[string]map[string][]map[string]string

type ApiTranslationsClient interface {
	AddToken(orgID, emailType, token string, translations []map[string]string) error
}

type Items struct {
	Items []*Item `json:"items" yaml:"items"`
}

type Item struct {
	Item  string `json:"item" yaml:"item"`
	Value string `json:"value" yaml:"value"`
}

func LoadTranslations(apiClient ApiTranslationsClient, c *config.Config) error {
	list, err := getTranslationsList(c)
	if err != nil {
		return err
	}

	totalKeys := len(list)
	currKey := 0
	for key, tokensList := range list {
		fmt.Println("============================")
		fmt.Println(fmt.Sprintf("Uploading keys %v / %v. Current: %s", currKey, totalKeys, key))

		totalTokens := len(tokensList)
		currToken := 0
		for tokenKey, translations := range tokensList {
			params := strings.Fields(key)
			orgID := params[0]
			emailType := params[1]

			fmt.Println("Upload translations: " + key + ". With token: " + tokenKey)
			err := apiClient.AddToken(orgID, emailType, tokenKey, translations)
			if err != nil {
				log.Fatalf("Can'emailType load trranslation: %s", err)
			} else {
				currToken++
			}
			fmt.Println(fmt.Sprintf("%v / %v", currToken, totalTokens))
		}
		currKey++
	}

	return nil
}

func getTranslationsList(c *config.Config) (translationsList, error) {
	data := make(map[string]map[string]map[string]map[string]string)
	res := make(map[string]map[string][]map[string]string)
	orgList := make(map[string]struct{})
	localeList := make(map[string]struct{})

	path := c.TranslationsPath
	locales, err := loadDir(path)
	if err != nil {
		return res, err
	}

	for _, l := range locales {
		locale := getLocale(l)
		localeDir := path + "/" + l
		files, err := loadDir(localeDir)
		if err != nil {
			return res, err
		}

		tokensPerType := map[string]map[string]map[string]string{}
		for _, file := range files {
			if isFileBlacklisted(c, file) {
				fmt.Println(fmt.Sprintf("Skipping, file %s blacklisted.", file))
				continue
			}

			filePath := localeDir + "/" + file
			content, err := ioutil.ReadFile(filePath)
			if err != nil {
				return res, err
			}

			unmarshalled := Items{}
			err = yaml.Unmarshal(content, &unmarshalled)
			if err != nil {
				return res, err
			}

			tokensInFile := make(map[string]map[string]string)
			for _, item := range unmarshalled.Items {
				translation := map[string]string{
					"locale":      locale,
					"translation": item.Value,
				}
				tokensInFile[item.Item] = translation
			}

			key := getTranslationsKey(c, file)
			tokensPerType[key] = tokensInFile
			orgList[key] = struct{}{}
		}
		localeList[locale] = struct{}{}
		data[locale] = tokensPerType
	}

	for orgKey, _ := range orgList {
		allTokensOnAllLangs := map[string][]map[string]string{}
		for locale, _ := range localeList {
			keys := data[locale]
			tokensInLang, ok := keys[orgKey]
			if !ok {
				continue
			}
			for token, v := range tokensInLang {
				allTokensOnAllLangs[token] = append(allTokensOnAllLangs[token], v)
			}
		}
		res[orgKey] = allTokensOnAllLangs
	}

	return res, nil
}

func loadDir(path string) ([]string, error) {
	var list []string
	dirs, err := filepath.Glob(path + "/*")
	if err != nil {
		return list, err
	}

	// Skip hidden folders
	for _, item := range dirs {
		if filepath.Base(item)[0] != '.' {
			l := filepath.Base(item)
			list = append(list, l)
		}
	}

	return list, nil
}

func getTranslationsKey(c *config.Config, filename string) string {
	orgID := uuid.UUID{}.String()
	if strings.Contains(filename, "sncf-") {
		orgID = c.OrgIDSNCF
		filename = strings.ReplaceAll(filename, "sncf-", "")
	}
	if strings.Contains(filename, "thalys-") {
		orgID = c.OrgIDThalys
		filename = strings.ReplaceAll(filename, "thalys-", "")
	}
	t := strings.ReplaceAll(filename, ".yaml", "")

	// tmp - Remove this if when enum naming will be fixed
	//if strings.Contains(t, "trip-cancelled") {
	//	t = strings.ReplaceAll(t, "trip-cancelled", "trip-canceled")
	//}

	return orgID + " " + t
}

func getLocale(locale string) string {
	return locale[0:2] + "-" + locale[2:]
}

func isFileBlacklisted(c *config.Config, filename string) bool {
	if strings.Contains(filename, "sncf-") && c.OrgIDSNCF == "" {
		fmt.Printf("SNCF orgID wasn't passed ")
		return true
	}

	if strings.Contains(filename, "thalys-") && c.OrgIDThalys == "" {
		fmt.Println("Thalys orgID wasn't passed ")
		return true
	}

	_, ok := blacklistedFiles[filename]

	return ok
}
