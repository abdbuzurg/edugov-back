package domain

import "time"

type InstitutionMagazine struct {
	ID            int64
	InstitutionID int64
	LanguageCode  string
	MagazineName  string
	Link          string
	LinkToRINC    string
	CreatedAt     time.Time
	UpdatedAt     time.Time
}
