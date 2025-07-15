package domain

import "time"

type InstitutionAccreditation struct {
	ID                int64
	InstitutionID     int64
	LanguageCode      string
	AccreditationType string
	GivenBy           string
	CreatedAt         time.Time
	UpdatedAt         time.Time
}
