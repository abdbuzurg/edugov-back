package domain

import "time"

type InstitutionPatent struct {
	ID               int64
	InstitutionID    int64
	LanguageCode     string
	PatentTitle      string
	Discipline       string
	Description      string
	ImplementedIn    string
	LinkToPatentFile string
	CreatedAt        time.Time
	UpdatedAt        time.Time
}
