package structs

type SoundCloudApiCreds struct {
	ClientId  string `json:"client_id"`
	AuthToken string `json:"auth_token"`
}

type Track struct {
	ArtworkUrl        string `json:"artwork_url"`
	Caption           string `json:"caption"`
	Commentable       bool   `json:"commentable"`
	CommentCount      int    `json:"comment_count"`
	CreatedAt         string `json:"created_at"`
	Description       string `json:"description"`
	Downloadable      bool   `json:"downloadable"`
	DownloadCount     int    `json:"download_count"`
	Duration          int64  `json:"duration"`
	FullDuration      int64  `json:"full_duration"`
	EmbeddableBy      string `json:"embeddable_by"`
	Genre             string `json:"genre"`
	HasDownloadsLeft  bool   `json:"has_downloads_left"`
	Id                int    `json:"id"`
	Kind              string `json:"kind"`
	LabelName         string `json:"label_name"`
	LastModified      string `json:"last_modified"`
	License           string `json:"license"`
	LikesCount        int    `json:"likes_count"`
	Permalink         string `json:"permalink"`
	PermalinkUrl      string `json:"permalink_url"`
	PlaybackCount     int    `json:"playback_count"`
	Public            bool   `json:"public"`
	PublisherMetadata struct {
		Id            int    `json:"id"`
		Urn           string `json:"urn"`
		ContainsMusic bool   `json:"contains_music"`
	} `json:"publisher_metadata"`
	PurchaseTitle string      `json:"purchase_title"`
	PurchaseUrl   string      `json:"purchase_url"`
	ReleaseDate   string      `json:"release_date"`
	RepostsCount  int         `json:"reposts_count"`
	SecretToken   string      `json:"secret_token"`
	Sharing       string      `json:"sharing"`
	State         string      `json:"state"`
	Streamable    bool        `json:"streamable"`
	TagList       string      `json:"tag_list"`
	Title         string      `json:"title"`
	TrackFormat   string      `json:"track_format"`
	Uri           string      `json:"uri"`
	Urn           string      `json:"urn"`
	UserId        int         `json:"user_id"`
	Visuals       interface{} `json:"visuals"`
	WaveformUrl   string      `json:"waveform_url"`
	DisplayDate   string      `json:"display_date"`
	Media         struct {
		Transcodings []struct {
			Url      string `json:"url"`
			Preset   string `json:"preset"`
			Duration int    `json:"duration"`
			Snipped  bool   `json:"snipped"`
			Format   struct {
				Protocol string `json:"protocol"`
				MimeType string `json:"mime_type"`
			} `json:"format"`
			Quality string `json:"quality"`
		} `json:"transcodings"`
	} `json:"media"`
	StationUrn         string `json:"station_urn"`
	StationPermalink   string `json:"station_permalink"`
	TrackAuthorization string `json:"track_authorization"`
	MonetizationModel  string `json:"monetization_model"`
	Policy             string `json:"policy"`
	User               struct {
		AvatarUrl      string `json:"avatar_url"`
		FirstName      string `json:"first_name"`
		FollowersCount int    `json:"followers_count"`
		FullName       string `json:"full_name"`
		Id             int    `json:"id"`
		Kind           string `json:"kind"`
		LastModified   string `json:"last_modified"`
		LastName       string `json:"last_name"`
		Permalink      string `json:"permalink"`
		PermalinkUrl   string `json:"permalink_url"`
		Uri            string `json:"uri"`
		Urn            string `json:"urn"`
		Username       string `json:"username"`
		Verified       bool   `json:"verified"`
		City           string `json:"city"`
		CountryCode    string `json:"country_code"`
		Badges         struct {
			Pro          bool `json:"pro"`
			ProUnlimited bool `json:"pro_unlimited"`
			Verified     bool `json:"verified"`
		} `json:"badges"`
		StationUrn       string `json:"station_urn"`
		StationPermalink string `json:"station_permalink"`
	} `json:"user"`
}

type PlayHistoryTrack struct {
	PlayedAt int64 `json:"played_at"`
	TrackId  int   `json:"track_id"`
	Track    Track `json:"track"`
}

type PlayHistory struct {
	Collection []struct {
		PlayHistoryTrack
	} `json:"collection"`
	NextHref string      `json:"next_href"`
	QueryUrn interface{} `json:"query_urn"`
}

type User struct {
	AvatarUrl             string `json:"avatar_url"`
	BlockedTracksCount    int    `json:"blocked_tracks_count"`
	City                  string `json:"city"`
	CommentsCount         int    `json:"comments_count"`
	ConsumerSubscriptions string `json:"consumer_subscriptions"`
	ConsumerSubscription  struct {
		Product struct {
			Id string `json:"id"`
		} `json:"product"`
	} `json:"consumer_subscription"`
	CountryCode          string `json:"country_code"`
	Cpp                  string `json:"cpp"`
	CreatedAt            string `json:"created_at"`
	CreatorSubscriptions []struct {
		Product struct {
			Id string `json:"id"`
		} `json:"product"`
	} `json:"creator_subscriptions"`
	CreatorSubscription struct {
		Product struct {
			Id string `json:"id"`
		} `json:"product"`
	} `json:"creator_subscription"`
	DateOfBirth struct {
		Month int `json:"month"`
		Year  int `json:"year"`
	} `json:"date_of_birth"`
	DefaultLicense          string `json:"default_license"`
	DefaultTracksFeedable   bool   `json:"default_tracks_feedable"`
	Description             string `json:"description"`
	DownloadsDisabled       bool   `json:"downloads_disabled"`
	DownloadsDisabledReason string `json:"downloads_disabled_reason"`
	FirstName               string `json:"first_name"`
	FollowersCount          int    `json:"followers_count"`
	FollowingsCount         int    `json:"followings_count"`
	FullName                string `json:"full_name"`
	Gender                  string `json:"gender"`
	GroupsCount             int    `json:"groups_count"`
	HiddenTracksCount       int    `json:"hidden_tracks_count"`
	Id                      int    `json:"id"`
	Kind                    string `json:"kind"`
	LastModified            string `json:"last_modified"`
	LastName                string `json:"last_name"`
	LikesCount              int    `json:"likes_count"`
	PlaylistLikesCount      int    `json:"playlist_likes_count"`
	Locale                  string `json:"locale"`
	Permalink               string `json:"permalink"`
	PermalinkUrl            string `json:"permalink_url"`
	PlaylistCount           int    `json:"playlist_count"`
	PrimaryEmail            string `json:"primary_email"`
	PrimaryEmailConfirmed   bool   `json:"primary_email_confirmed"`
	PrimaryEmailSha256      string `json:"primary_email_sha256"`
	PrivatePlaylistsCount   int    `json:"private_playlists_count"`
	PrivateTracksCount      int    `json:"private_tracks_count"`
	Quota                   struct {
		UnlimitedUploadQuota         bool `json:"unlimited_upload_quota"`
		UploadSecondsUsed            int  `json:"upload_seconds_used"`
		UploadSecondsLeft            int  `json:"upload_seconds_left"`
		UploadTracksUsed             int  `json:"upload_tracks_used"`
		UnlimitedUploadDurationQuota bool `json:"unlimited_upload_duration_quota"`
		UnlimitedUploadTrackQuota    bool `json:"unlimited_upload_track_quota"`
	} `json:"quota"`
	RepostsCount int    `json:"reposts_count"`
	TrackCount   int    `json:"track_count"`
	Urn          string `json:"urn"`
	Uri          string `json:"uri"`
	Username     string `json:"username"`
	Verified     bool   `json:"verified"`
	Visuals      struct {
		Urn     string `json:"urn"`
		Enabled bool   `json:"enabled"`
		Visuals []struct {
			Urn       string `json:"urn"`
			EntryTime int    `json:"entry_time"`
			VisualUrl string `json:"visual_url"`
		} `json:"visuals"`
		Tracking interface{} `json:"tracking"`
	} `json:"visuals"`
	Confirmed bool `json:"confirmed"`
	Badges    struct {
		Pro          bool `json:"pro"`
		ProUnlimited bool `json:"pro_unlimited"`
		Verified     bool `json:"verified"`
	} `json:"badges"`
	AnalyticsId          string `json:"analytics_id"`
	ConsentManagementJwt struct {
		UserId string `json:"userId"`
		Jwt    string `json:"jwt"`
	} `json:"consent_management_jwt"`
	StationUrn       string `json:"station_urn"`
	StationPermalink string `json:"station_permalink"`
	MarketingIds     struct {
		Gtm string `json:"gtm"`
	} `json:"marketing_ids"`
}
