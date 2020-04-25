package reddit

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"

	"../model"
)

const (
	apiAccessTokenURL = "https://www.reddit.com/api/v1/access_token"
)

// NewBotFromFile initializes bot from local file
func NewBotFromFile(file string) (*model.Bot, error) {
	var cred model.APICredentials

	bot := &model.Bot{}

	data, err := ioutil.ReadFile(fmt.Sprintf(file))
	if err != nil {
		return bot, err
	}

	err = json.Unmarshal(data, &cred)
	if err != nil {
		return bot, err
	}

	auth, err := authorizeAccount(cred)
	if err != nil {
		return bot, err
	}

	bot = &model.Bot{
		Account:        cred,
		Authentication: auth,
	}

	return bot, nil

}

func authorizeAccount(credentials model.APICredentials) (model.APITokenCredentials, error) {
	client := &http.Client{}
	requestBody := url.Values{}
	creds := model.APITokenCredentials{}

	requestBody.Add("grant_type", "password")
	requestBody.Add("username", credentials.AccountUsername)
	requestBody.Add("password", credentials.AccountPassword)
	requestBody.Add("duration", "permanent")

	r, err := http.NewRequest("POST", apiAccessTokenURL, strings.NewReader(requestBody.Encode()))
	if err != nil {
		return creds, err
	}

	r.SetBasicAuth(credentials.ClientID, credentials.ClientSecret)

	r.Header.Set("User-Agent", credentials.UserAgent)
	r.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	resp, err := client.Do(r)
	if err != nil {
		return creds, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return creds, err
	}

	err = json.Unmarshal(body, &creds)

	if err != nil {
		return creds, err
	}

	return creds, nil
}
