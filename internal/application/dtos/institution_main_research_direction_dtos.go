package dtos

import "time"

// ---- REQUEST DTOs ----

type CreateInstitutionMainResearchDirectionRequest struct {
	InstitutionID          int64  `json:"institutionID" validate:"required,min=1"`
	LanguageCode           string `json:"-" validate:"required,len=2"`
	ResearchDirectionTitle string `json:"researchDirectionTitle" validate:"required"`
	Discipline             string `json:"discipline" validate:"required"`
	AreaOfResearch         string `json:"areaOfResearch" validate:"required"`
}

type UpdateInstitutionMainResearchDirectionRequest struct {
	ID                     int64   `json:"id" validate:"required,min=1"`
	ResearchDirectionTitle *string `json:"researchDirectionTitle" validate:"omitempty"`
	Discipline             *string `json:"discipline" validate:"omitempty"`
	AreaOfResearch         *string `json:"areaOfResearch" validate:"omitempty"`
}

// ---- RESPONSE DTOs ----

type InstitutionMainResearchDirectionResponse struct {
	ID                     int64     `json:"id"`
	ResearchDirectionTitle string    `json:"researchDirectionTitle"`
	Discipline             string    `json:"discipline"`
	AreaOfResearch         string    `json:"areaOfResearch"`
	CreatedAt              time.Time `json:"createdAt"`
	UpdatedAt              time.Time `json:"updatedAt"`
}
