package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

const projectID = "email-templates"

// Decimeter: |
const d = "%7C"

type APIClient struct {
	apiPath string
	apiKey  string
}

func Init(apiKey, apiHost string) *APIClient {
	c := &APIClient{}
	c.apiPath = apiHost + "/v1/translations/" + projectID + d
	c.apiKey = apiKey

	return c
}

func (a *APIClient) AddToken(orgID, emailType, token string, translations []map[string]string) error {
	uri := a.getURI(orgID, emailType, token)
	fmt.Println(uri)
	data := map[string][]map[string]string{
		"translations": translations,
	}
	j, err := json.Marshal(data)
	if err != nil {
		return err
	}

	client := &http.Client{}
	request, err := http.NewRequest(http.MethodPut, uri, bytes.NewBuffer(j))
	if err != nil {
		return err
	}
	request.Header.Add("Authorization", "Bearer "+a.apiKey)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("Accept", "application/json")

	response, err := client.Do(request)
	if err != nil {
		return err
	}
	if response.StatusCode != http.StatusOK {
		fmt.Println("Error during upload translation")
		fmt.Println(response)
	}
	defer response.Body.Close()

	return nil
}

func (a *APIClient) getURI(orgID, emailType, token string) string {
	return a.apiPath + orgID + d + emailType + d + token
}
