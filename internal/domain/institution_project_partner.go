package domain

import "time"

type InstitutionProjectPartner struct {
	ID                   int64
	InstitutionProjectID int64
	PartnerType          string
	PartnerName          string
	LinkToPartner        string
	LanguageCode         string
	CreatedAt            time.Time
	UpdatedAt            time.Time
}
