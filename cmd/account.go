package cmd

import (
	"encoding/json"

	"../model"
)

const (
	me             = "https://oauth.reddit.com/api/v1/me"
	blocked        = "https://oauth.reddit.com/api/v1/me/blocked"
	friends        = "https://oauth.reddit.com/api/v1/me/friends"
	karma          = "https://oauth.reddit.com/api/v1/me/karma"
	prefs          = "https://oauth.reddit.com/api/v1/me/prefs"
	trophies       = "https://oauth.reddit.com/api/v1/me/trophies"
	prefsBlocked   = "https://oauth.reddit.com/prefs/blocked"
	prefsFriends   = "https://oauth.reddit.com/prefs/friends"
	prefsMessaging = "https://oauth.reddit.com/prefs/messaging"
	prefsTrusted   = "https://oauth.reddit.com/prefs/trusted"
	prefsWhere     = "https://oauth.reddit.com/prefs/where"
)

// Me targets endpoint: /api/v1/me
func Me(bot *model.Bot) (model.APIv1Me, error) {
	res := model.APIv1Me{}
	resp, err := Bodyless(bot, "GET", "https://oauth.reddit.com/api/v1/me")
	if err != nil {
		return res, err
	}

	err = json.Unmarshal(resp, &res)
	if err != nil {
		return res, err
	}
	return res, nil
}

// Karma targets endpoint: /api/v1/me/karma
func Karma(bot *model.Bot) (model.APIv1MeKarma, error) {
	res := model.APIv1MeKarma{}
	resp, err := Bodyless(bot, "GET", "https://oauth.reddit.com/api/v1/me/karma")
	if err != nil {
		return res, err
	}

	err = json.Unmarshal(resp, &res)
	if err != nil {
		return res, err
	}
	return res, nil
}

// Prefs targets endpoint: /api/v1/me/prefs
func Prefs(bot *model.Bot) (model.APIv1MePrefs, error) {
	res := model.APIv1MePrefs{}
	resp, err := Bodyless(bot, "GET", "https://oauth.reddit.com/api/v1/me/prefs")
	if err != nil {
		return res, err
	}

	err = json.Unmarshal(resp, &res)
	if err != nil {
		return res, err
	}
	return res, nil
}

// Trophies targets endpoint: /api/v1/me/trophies
func Trophies(bot *model.Bot) (model.APIv1MeTrophies, error) {
	res := model.APIv1MeTrophies{}
	resp, err := Bodyless(bot, "GET", "https://oauth.reddit.com/api/v1/me/trophies")
	if err != nil {
		return res, err
	}

	err = json.Unmarshal(resp, &res)
	if err != nil {
		return res, err
	}

	return res, nil
}
