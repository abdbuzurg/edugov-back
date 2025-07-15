package domain

import "time"

type InstitutionResearchSupportInfrastructure struct {
	ID                                 int64
	InstitutionID                      int64
	LanguageCode                       string
	ResearchSupportInfrastructureTitle string
	ResearchSupportInfrastructureType  string
	TINOfLegalEntity                   string
	CreatedAt                          time.Time
	UpdatedAt                          time.Time
}
