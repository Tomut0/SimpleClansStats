package domain

import "time"

type Server struct {
	ID           string
	DisplayName  string
	Disabled     bool
	Online       bool
	LastIngestAt *time.Time
}

type ServerRecord struct {
	ID, DisplayName, LastReceivedAt string
	Disabled                        bool
	OnlineTimeoutMinutes            int64
}

type ServerRegistration struct {
	Server Server
	APIKey string
}
