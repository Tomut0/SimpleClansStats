package domain

type AppSettings struct {
	AppName                                         string
	RemoveAuthor                                    bool
	RemoveTelemetry                                 bool
	PrimaryLanguage                                 string
	Multilingual                                    bool
	TimeZone, DateFormat, TimeFormat                string
	PlayerAvatarURL                                 string
	Theme                                           string
	LogoVersion                                     int64
	ServerOnlineTimeoutMinutes, IngestRetentionDays int64
}

type AppSettingsUpdate struct {
	AppName                                         string
	RemoveAuthor                                    bool
	RemoveTelemetry                                 bool
	PrimaryLanguage                                 string
	Multilingual                                    bool
	TimeZone, DateFormat, TimeFormat                string
	PlayerAvatarURL                                 string
	ServerOnlineTimeoutMinutes, IngestRetentionDays *int64
	Logo                                            []byte
}
