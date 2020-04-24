package reddit

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"

	"../model"
)

const (
	apiAccessTokenURL = "https://www.reddit.com/api/v1/access_token"
)

// NewBotFromFile initializes bot from local file
func NewBotFromFile(file string) (model.APICredentials, model.APITokenCredentials) {
	var cred model.APICredentials

	data, err := ioutil.ReadFile(fmt.Sprintf(file))
	if err != nil {
		log.Fatalf("\nError loading credential file:\nError: %s", err)
	}

	err = json.Unmarshal(data, &cred)
	if err != nil {
		log.Fatalf("\nError unpacking JSON credential file:\nError: %s\nFile contents: %s", err, data)
	}

	return cred, authorizeAccount(cred)

}

func authorizeAccount(credentials model.APICredentials) model.APITokenCredentials {
	client := &http.Client{}
	requestBody := url.Values{}

	requestBody.Add("grant_type", "client_credentials")
	requestBody.Add("username", credentials.AccountUsername)
	requestBody.Add("password", credentials.AccountPassword)
	requestBody.Add("duration", "permanent")

	r, err := http.NewRequest("POST", apiAccessTokenURL, strings.NewReader(requestBody.Encode()))
	if err != nil {
		log.Fatal(err)
	}

	r.SetBasicAuth(credentials.ClientID, credentials.ClientSecret)

	r.Header.Set("User-Agent", credentials.UserAgent)
	r.Header.Add("Content-Type", "application/x-www-form-urlencoded; param=value")

	resp, err := client.Do(r)
	if err != nil {
		log.Fatal(err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	creds := model.APITokenCredentials{}

	err = json.Unmarshal(body, &creds)

	if err != nil {
		log.Fatal(err)
	}

	return creds
}
