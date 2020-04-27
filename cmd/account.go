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

// Blocked targets endpoint: /prefs/blocked
func Blocked(bot *model.Bot) (model.PrefsBlocked, error) {
	res := model.PrefsBlocked{}
	resp, err := Bodyless(bot, "GET", "https://oauth.reddit.com/prefs/blocked")
	if err != nil {
		return res, err
	}

	err = json.Unmarshal(resp, &res)
	if err != nil {
		return res, err
	}
	return res, nil
}

// Friends targets endpoint: /prefs/friends
func Friends(bot *model.Bot) (model.PrefsFriends, error) {
	res := model.PrefsFriends{}
	resp, err := Bodyless(bot, "GET", "https://oauth.reddit.com/prefs/friends")
	if err != nil {
		return res, err
	}

	err = json.Unmarshal(resp, &res)
	if err != nil {
		return res, err
	}
	return res, nil
}

// Messaging targets endpoint: /prefs/messaging
func Messaging(bot *model.Bot) (model.PrefsMessaging, error) {
	res := model.PrefsMessaging{}
	resp, err := Bodyless(bot, "GET", "https://oauth.reddit.com/prefs/messaging")
	if err != nil {
		return res, err
	}

	err = json.Unmarshal(resp, &res)
	if err != nil {
		return res, err
	}
	return res, nil
}

// Trusted targets endpoint: /prefs/trusted
func Trusted(bot *model.Bot) (model.PrefsTrusted, error) {
	res := model.PrefsTrusted{}
	resp, err := Bodyless(bot, "GET", "https://oauth.reddit.com/prefs/trusted")
	if err != nil {
		return res, err
	}

	err = json.Unmarshal(resp, &res)
	if err != nil {
		return res, err
	}
	return res, nil
}

// Where targets endpoint: /prefs/where
func Where(bot *model.Bot) (model.PrefsBlocked, error) {
	res := model.PrefsBlocked{}
	resp, err := Bodyless(bot, "GET", "https://oauth.reddit.com/prefs/where")
	if err != nil {
		return res, err
	}

	err = json.Unmarshal(resp, &res)
	if err != nil {
		return res, err
	}
	return res, nil
}

// CollectSelfInfo returns chunk of info
func CollectSelfInfo(bot *model.Bot) (model.CollectSelfInfo, error) {
	alpha, err := Me(bot)
	beta, err := Karma(bot)
	charlie, err := Prefs(bot)
	delta, err := Trophies(bot)
	echo, err := Blocked(bot)
	foxtrox, err := Friends(bot)
	gamma, err := Messaging(bot)
	hotel, err := Trusted(bot)

	self := model.CollectSelfInfo{
		Me:        alpha,
		Karma:     beta,
		Prefs:     charlie,
		Trophies:  delta,
		Blocked:   echo,
		Friends:   foxtrox,
		Messaging: gamma,
		Trusted:   hotel,
	}

	return self, err
}
