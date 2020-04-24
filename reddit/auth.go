package reddit

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
)

const (
	apiAccessTokenURL = "https://www.reddit.com/api/v1/access_token"
)

type apiCredentials struct {
	UserAgent       string `json:"user_agent"`
	AccountUsername string `json:"username"`
	AccountPassword string `json:"password"`
	ClientID        string `json:"client_id"`
	ClientSecret    string `json:"client_secret"`
}

type apiTokenCredentials struct {
	AccessToken  string `json:"access_token"`
	TokenType    string `json:"token_type"`
	ExpiresIn    int    `json:"expires_in"`
	RefreshToken string `json:"refresh_token"`
	Scope        string `json:"scope"`
}

func newBotFromFile(file string) apiTokenCredentials {
	var cred apiCredentials

	data, _ := ioutil.ReadFile(fmt.Sprintf("../%s", file))

	err := json.Unmarshal(data, &cred)
	if err != nil {
		log.Fatal(err)
	}

	delta := authorizeAccount(cred)

	return delta

}

func authorizeAccount(credentials apiCredentials) apiTokenCredentials {
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

	creds := apiTokenCredentials{}

	err = json.Unmarshal(body, &creds)
	if err != nil {
		log.Fatal(err)
	}

	return creds
}
