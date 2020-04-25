package cmd

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"../model"
)

// Bodyless HTTP GET request
func Bodyless(bot *model.Bot, method string, url string) ([]byte, error) {
	credentials := &bot.Account
	api := &bot.Authentication
	client := &http.Client{}

	r, err := http.NewRequest(method, url, nil)
	if err != nil {
		return nil, err
	}

	r.Header.Set("User-Agent", credentials.UserAgent)
	r.Header.Add("Authorization", fmt.Sprintf("bearer %s", api.AccessToken))

	resp, err := client.Do(r)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("Bad status code: %v", resp.StatusCode)
	}

	return body, nil

}
