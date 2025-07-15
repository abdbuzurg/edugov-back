package domain

import "time"

type InstitutionDetails struct {
	ID               int64
	InstitutionID     int64
	LanguageCode     string
	InstitutionTitle string
	InstitutionType  string
	LegalStatus      string
	Mission          string
	Founder          string
	LegalAddress     string
	FactualAddress   *string
	CreatedAt        time.Time
	UpdatedAt        time.Time
}
