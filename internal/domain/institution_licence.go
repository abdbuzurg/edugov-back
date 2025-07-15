package domain

import "time"

type InstitutionLicence struct {
	ID            int64
	InstitutionID int64
	LanguageCode  string
	LicenceTitle  string
	LicenceType   string
	GivenBy       string
	LinkToFile    string
	DateStart      time.Time
	DateEnd        time.Time
	CreatedAt     time.Time
	UpdatedAt     time.Time
}
