package domain

import "time"

type InstitutionProject struct {
	ID              int64
	InstitutionID   int64
	LanguageCode    string
	ProjectType     string
	ProjectTitle    string
	DateStart       time.Time
	DateEnd         time.Time
	Fund            float64
	InstitutionRole string
	Coordinator     string
	Partners        []*InstitutionProjectPartner
	CreatedAt       time.Time
	UpdatedAt       time.Time
}
