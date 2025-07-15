package dtos

import "time"

// ---- REQUEST DTOs ----

type CreateInstitutionResearchSupportInfrastructureRequest struct {
	InstitutionID                      int64  `json:"institutionID" validate:"required,min=1"`
	LanguageCode                       string `json:"-" validate:"required,len=2"`
	ResearchSupportInfrastructureTitle string `json:"researchSupportInfrastructureTitle" validate:"required"`
	ResearchSupportInfrastructureType  string `json:"researchSupportInfrastructureType" validate:"required"`
	TINOfLegalEntity                   string `json:"tinOfLegalEntity" validate:"required"`
}

type UpdateInstitutionResearchSupportInfrastructureRequest struct {
	ID                                 int64   `json:"id" validate:"required,min=1"`
	ResearchSupportInfrastructureTitle *string `json:"researchSupportInfrastructureTitle" validate:"omitempty"`
	ResearchSupportInfrastructureType  *string `json:"researchSupportInfrastructureType" validate:"omitempty"`
	TINOfLegalEntity                   *string `json:"tinOfLegalEntity" validate:"omitempty"`
}

// ---- RESPONSE DTOs ----

type InstitutionResearchSupportInfrastructureResponse struct {
	ID                                 int64     `json:"id"`
	ResearchSupportInfrastructureTitle string    `json:"researchSupportInfrastructureTitle"`
	ResearchSupportInfrastructureType  string    `json:"researchSupportInfrastructureType"`
	TINOfLegalEntity                   string    `json:"tinOfLegalEntity"`
	CreatedAt                          time.Time `json:"createdAt"`
	UpdatedAt                          time.Time `json:"updatedAt"`
}
