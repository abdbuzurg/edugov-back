package domain

import "time"

type InstitutionConference struct {
	ID               int64
	InstitutionID    int64
	LanguageCode     string
	ConferenceTitle  string
	Link             string
	LinkToRINC       string
	DateOfConference time.Time
	CreatedAt        time.Time
	UpdatedAt        time.Time
}
