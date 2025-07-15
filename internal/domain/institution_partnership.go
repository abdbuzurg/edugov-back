package domain

import "time"

type InstitutionPartnership struct {
	ID             int64
	InstitutionID  int64
	LanguageCode   string
	PartnerName    string
	PartnerType    string
	Goal           string
	LinkToPartner  string
	DateOfContract time.Time
	CreatedAt      time.Time
	UpdatedAt      time.Time
}
