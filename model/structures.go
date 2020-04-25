package model

// APICredentials holds the information from user's .cred file
type APICredentials struct {
	UserAgent       string `json:"user_agent"`
	AccountUsername string `json:"username"`
	AccountPassword string `json:"password"`
	ClientID        string `json:"client_id"`
	ClientSecret    string `json:"client_secret"`
}

// APITokenCredentials return struct to give AccessToken, TokenType, etc
type APITokenCredentials struct {
	AccessToken  string `json:"access_token"`
	TokenType    string `json:"token_type"`
	ExpiresIn    int    `json:"expires_in"`
	RefreshToken string `json:"refresh_token"`
	Scope        string `json:"scope"`
}

// Bot Joins the two
type Bot struct {
	Account        APICredentials
	Authentication APITokenCredentials
}

// APIv1Me endpoint: https://oauth.reddit.com/api/v1/me
type APIv1Me struct {
	IsEmployee           bool   `json:"is_employee"`
	SeenLayoutSwitch     bool   `json:"seen_layout_switch"`
	HasVisitedNewProfile bool   `json:"has_visited_new_profile"`
	PrefNoProfanity      bool   `json:"pref_no_profanity"`
	HasExternalAccount   bool   `json:"has_external_account"`
	PrefGeopopular       string `json:"pref_geopopular"`
	SeenRedesignModal    bool   `json:"seen_redesign_modal"`
	PrefShowTrending     bool   `json:"pref_show_trending"`
	Subreddit            struct {
		DefaultSet                 bool          `json:"default_set"`
		UserIsContributor          bool          `json:"user_is_contributor"`
		BannerImg                  string        `json:"banner_img"`
		RestrictPosting            bool          `json:"restrict_posting"`
		UserIsBanned               bool          `json:"user_is_banned"`
		FreeFormReports            bool          `json:"free_form_reports"`
		CommunityIcon              string        `json:"community_icon"`
		ShowMedia                  bool          `json:"show_media"`
		IconColor                  string        `json:"icon_color"`
		UserIsMuted                bool          `json:"user_is_muted"`
		DisplayName                string        `json:"display_name"`
		HeaderImg                  interface{}   `json:"header_img"`
		Title                      string        `json:"title"`
		Coins                      int           `json:"coins"`
		PreviousNames              []interface{} `json:"previous_names"`
		Over18                     bool          `json:"over_18"`
		IconSize                   []int         `json:"icon_size"`
		PrimaryColor               string        `json:"primary_color"`
		IconImg                    string        `json:"icon_img"`
		Description                string        `json:"description"`
		SubmitLinkLabel            string        `json:"submit_link_label"`
		HeaderSize                 interface{}   `json:"header_size"`
		RestrictCommenting         bool          `json:"restrict_commenting"`
		Subscribers                int           `json:"subscribers"`
		SubmitTextLabel            string        `json:"submit_text_label"`
		IsDefaultIcon              bool          `json:"is_default_icon"`
		LinkFlairPosition          string        `json:"link_flair_position"`
		DisplayNamePrefixed        string        `json:"display_name_prefixed"`
		KeyColor                   string        `json:"key_color"`
		Name                       string        `json:"name"`
		IsDefaultBanner            bool          `json:"is_default_banner"`
		URL                        string        `json:"url"`
		BannerSize                 interface{}   `json:"banner_size"`
		UserIsModerator            bool          `json:"user_is_moderator"`
		PublicDescription          string        `json:"public_description"`
		LinkFlairEnabled           bool          `json:"link_flair_enabled"`
		DisableContributorRequests bool          `json:"disable_contributor_requests"`
		SubredditType              string        `json:"subreddit_type"`
		UserIsSubscriber           bool          `json:"user_is_subscriber"`
	} `json:"subreddit"`
	IsSponsor           bool        `json:"is_sponsor"`
	GoldExpiration      interface{} `json:"gold_expiration"`
	HasGoldSubscription bool        `json:"has_gold_subscription"`
	NumFriends          int         `json:"num_friends"`
	Features            struct {
		PromotedTrendBlanks       bool `json:"promoted_trend_blanks"`
		ShowAmpLink               bool `json:"show_amp_link"`
		Chat                      bool `json:"chat"`
		TwitterEmbed              bool `json:"twitter_embed"`
		IsEmailPermissionRequired bool `json:"is_email_permission_required"`
		ModAwards                 bool `json:"mod_awards"`
		MwebXpromoRevampV3        struct {
			Owner        string `json:"owner"`
			Variant      string `json:"variant"`
			ExperimentID int    `json:"experiment_id"`
		} `json:"mweb_xpromo_revamp_v3"`
		ChatSubreddit                                      bool `json:"chat_subreddit"`
		AwardsOnStreams                                    bool `json:"awards_on_streams"`
		MwebXpromoModalListingClickDailyDismissibleIos     bool `json:"mweb_xpromo_modal_listing_click_daily_dismissible_ios"`
		CommunityAwards                                    bool `json:"community_awards"`
		ModlogCopyrightRemoval                             bool `json:"modlog_copyright_removal"`
		DoNotTrack                                         bool `json:"do_not_track"`
		ChatUserSettings                                   bool `json:"chat_user_settings"`
		MwebXpromoInterstitialCommentsIos                  bool `json:"mweb_xpromo_interstitial_comments_ios"`
		NoreferrerToNoopener                               bool `json:"noreferrer_to_noopener"`
		PremiumSubscriptionsTable                          bool `json:"premium_subscriptions_table"`
		MwebXpromoInterstitialCommentsAndroid              bool `json:"mweb_xpromo_interstitial_comments_android"`
		MwebXpromoModalListingClickDailyDismissibleAndroid bool `json:"mweb_xpromo_modal_listing_click_daily_dismissible_android"`
		StreamAsAPostType                                  bool `json:"stream_as_a_post_type"`
		ChatGroupRollout                                   bool `json:"chat_group_rollout"`
		CustomFeeds                                        bool `json:"custom_feeds"`
		SpezModal                                          bool `json:"spez_modal"`
		MwebSharingClipboard                               struct {
			Owner        string `json:"owner"`
			Variant      string `json:"variant"`
			ExperimentID int    `json:"experiment_id"`
		} `json:"mweb_sharing_clipboard"`
		ExpensiveCoinsPackage bool `json:"expensive_coins_package"`
	} `json:"features"`
	HasAndroidSubscription  bool          `json:"has_android_subscription"`
	Verified                bool          `json:"verified"`
	NewModmailExists        bool          `json:"new_modmail_exists"`
	PrefAutoplay            bool          `json:"pref_autoplay"`
	Coins                   int           `json:"coins"`
	HasPaypalSubscription   bool          `json:"has_paypal_subscription"`
	HasSubscribedToPremium  bool          `json:"has_subscribed_to_premium"`
	ID                      string        `json:"id"`
	HasStripeSubscription   bool          `json:"has_stripe_subscription"`
	SeenPremiumAdblockModal bool          `json:"seen_premium_adblock_modal"`
	CanCreateSubreddit      bool          `json:"can_create_subreddit"`
	Over18                  bool          `json:"over_18"`
	IsGold                  bool          `json:"is_gold"`
	IsMod                   bool          `json:"is_mod"`
	SuspensionExpirationUtc interface{}   `json:"suspension_expiration_utc"`
	HasVerifiedEmail        bool          `json:"has_verified_email"`
	IsSuspended             bool          `json:"is_suspended"`
	PrefVideoAutoplay       bool          `json:"pref_video_autoplay"`
	InChat                  bool          `json:"in_chat"`
	CanEditName             bool          `json:"can_edit_name"`
	InRedesignBeta          bool          `json:"in_redesign_beta"`
	IconImg                 string        `json:"icon_img"`
	HasModMail              bool          `json:"has_mod_mail"`
	PrefNightmode           bool          `json:"pref_nightmode"`
	OauthClientID           string        `json:"oauth_client_id"`
	HideFromRobots          bool          `json:"hide_from_robots"`
	PasswordSet             bool          `json:"password_set"`
	LinkKarma               int           `json:"link_karma"`
	ForcePasswordReset      bool          `json:"force_password_reset"`
	SeenGiveAwardTooltip    bool          `json:"seen_give_award_tooltip"`
	InboxCount              int           `json:"inbox_count"`
	PrefTopKarmaSubreddits  bool          `json:"pref_top_karma_subreddits"`
	HasMail                 bool          `json:"has_mail"`
	PrefShowSnoovatar       bool          `json:"pref_show_snoovatar"`
	Name                    string        `json:"name"`
	PrefClickgadget         int           `json:"pref_clickgadget"`
	Created                 float64       `json:"created"`
	GoldCreddits            int           `json:"gold_creddits"`
	CreatedUtc              float64       `json:"created_utc"`
	HasIosSubscription      bool          `json:"has_ios_subscription"`
	PrefShowTwitter         bool          `json:"pref_show_twitter"`
	InBeta                  bool          `json:"in_beta"`
	CommentKarma            int           `json:"comment_karma"`
	HasSubscribed           bool          `json:"has_subscribed"`
	LinkedIdentities        []interface{} `json:"linked_identities"`
	SeenSubredditChatFtux   bool          `json:"seen_subreddit_chat_ftux"`
}

// APIv1MeKarma endpoint: https://oauth.reddit.com/api/v1/me/karma
type APIv1MeKarma struct {
	Kind string `json:"kind"`
	Data []struct {
		Sr           string `json:"sr"`
		CommentKarma int    `json:"comment_karma"`
		LinkKarma    int    `json:"link_karma"`
	} `json:"data"`
}

// APIv1MePrefs endpoint: https://oauth.reddit.com/api/v1/me/prefs
type APIv1MePrefs struct {
	DefaultThemeSr                        interface{} `json:"default_theme_sr"`
	ThreadedMessages                      bool        `json:"threaded_messages"`
	HideDowns                             bool        `json:"hide_downs"`
	LabelNsfw                             bool        `json:"label_nsfw"`
	ActivityRelevantAds                   bool        `json:"activity_relevant_ads"`
	EmailMessages                         bool        `json:"email_messages"`
	ProfileOptOut                         bool        `json:"profile_opt_out"`
	VideoAutoplay                         bool        `json:"video_autoplay"`
	ThirdPartySiteDataPersonalizedContent bool        `json:"third_party_site_data_personalized_content"`
	ShowLinkFlair                         bool        `json:"show_link_flair"`
	ShowTrending                          bool        `json:"show_trending"`
	PrivateFeeds                          bool        `json:"private_feeds"`
	MonitorMentions                       bool        `json:"monitor_mentions"`
	PublicServerSeconds                   bool        `json:"public_server_seconds"`
	Research                              bool        `json:"research"`
	IgnoreSuggestedSort                   bool        `json:"ignore_suggested_sort"`
	SendCrosspostMessages                 bool        `json:"send_crosspost_messages"`
	EmailDigests                          bool        `json:"email_digests"`
	Layout                                int         `json:"layout"`
	NumComments                           int         `json:"num_comments"`
	Clickgadget                           bool        `json:"clickgadget"`
	UseGlobalDefaults                     bool        `json:"use_global_defaults"`
	ShowSnoovatar                         bool        `json:"show_snoovatar"`
	Over18                                bool        `json:"over_18"`
	ShowStylesheets                       bool        `json:"show_stylesheets"`
	LiveOrangereds                        bool        `json:"live_orangereds"`
	EnableDefaultThemes                   bool        `json:"enable_default_themes"`
	LegacySearch                          bool        `json:"legacy_search"`
	DomainDetails                         bool        `json:"domain_details"`
	CollapseLeftBar                       bool        `json:"collapse_left_bar"`
	Lang                                  string      `json:"lang"`
	HideUps                               bool        `json:"hide_ups"`
	ThirdPartyDataPersonalizedAds         bool        `json:"third_party_data_personalized_ads"`
	AllowClicktracking                    bool        `json:"allow_clicktracking"`
	HideFromRobots                        bool        `json:"hide_from_robots"`
	ShowTwitter                           bool        `json:"show_twitter"`
	Compress                              bool        `json:"compress"`
	StoreVisits                           bool        `json:"store_visits"`
	ThreadedModmail                       bool        `json:"threaded_modmail"`
	DesignBeta                            bool        `json:"design_beta"`
	MinLinkScore                          int         `json:"min_link_score"`
	MediaPreview                          string      `json:"media_preview"`
	ShowLocationBasedRecommendations      bool        `json:"show_location_based_recommendations"`
	Nightmode                             bool        `json:"nightmode"`
	HighlightControversial                bool        `json:"highlight_controversial"`
	Geopopular                            string      `json:"geopopular"`
	ThirdPartySiteDataPersonalizedAds     bool        `json:"third_party_site_data_personalized_ads"`
	SurveyLastSeenTime                    interface{} `json:"survey_last_seen_time"`
	MinCommentScore                       int         `json:"min_comment_score"`
	PublicVotes                           bool        `json:"public_votes"`
	CollapseReadMessages                  bool        `json:"collapse_read_messages"`
	ShowFlair                             bool        `json:"show_flair"`
	MarkMessagesRead                      bool        `json:"mark_messages_read"`
	SearchIncludeOver18                   bool        `json:"search_include_over_18"`
	NoProfanity                           bool        `json:"no_profanity"`
	HideAds                               bool        `json:"hide_ads"`
	Beta                                  bool        `json:"beta"`
	TopKarmaSubreddits                    bool        `json:"top_karma_subreddits"`
	Newwindow                             bool        `json:"newwindow"`
	Numsites                              int         `json:"numsites"`
	Media                                 string      `json:"media"`
	SendWelcomeMessages                   bool        `json:"send_welcome_messages"`
	ShowGoldExpiration                    bool        `json:"show_gold_expiration"`
	HighlightNewComments                  bool        `json:"highlight_new_comments"`
	EmailUnsubscribeAll                   bool        `json:"email_unsubscribe_all"`
	DefaultCommentSort                    string      `json:"default_comment_sort"`
	AcceptPms                             string      `json:"accept_pms"`
}

// APIv1MeTrophies endpoint: https://oauth.reddit.com/api/v1/me/trophies
type APIv1MeTrophies struct {
	Kind string `json:"kind"`
	Data struct {
		Trophies []struct {
			Kind string `json:"kind"`
			Data struct {
				Icon70      string      `json:"icon_70"`
				Name        string      `json:"name"`
				URL         interface{} `json:"url"`
				Icon40      string      `json:"icon_40"`
				AwardID     string      `json:"award_id"`
				ID          string      `json:"id"`
				Description interface{} `json:"description"`
			} `json:"data"`
		} `json:"trophies"`
	} `json:"data"`
}
