package domain

import "time"

type InstitutionMainResearchDirection struct {
	ID                      int64
	InstitutionID           int64
	LanguageCode            string
	ResearchDirectionTitle string
	Discipline              string
	AreaOfResearch          string
	CreatedAt               time.Time
	UpdatedAt               time.Time
}
