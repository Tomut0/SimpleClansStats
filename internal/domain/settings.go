package domain

type AppSettings struct {
	AppName                                         string
	RemoveAuthor                                    bool
	TimeZone, DateFormat, TimeFormat                string
	PlayerAvatarURL                                 string
	LogoVersion                                     int64
	ServerOnlineTimeoutMinutes, IngestRetentionDays int64
}

type AppSettingsUpdate struct {
	AppName                                         string
	RemoveAuthor                                    bool
	TimeZone, DateFormat, TimeFormat                string
	PlayerAvatarURL                                 string
	ServerOnlineTimeoutMinutes, IngestRetentionDays *int64
	Logo                                            []byte
}
