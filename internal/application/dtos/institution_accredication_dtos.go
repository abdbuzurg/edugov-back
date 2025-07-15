package dtos

import "time"

// ---- REQUEST DTOs ----

type CreateInstitutionAccreditationRequest struct {
	InstitutionID     int64  `json:"institutionID" validate:"required,min=1"`
	LanguageCode      string `json:"-" validate:"required,len=2"`
	AccreditationType string `json:"accreditationType" validate:"required"`
	GivenBy           string `json:"givenBy" validate:"required"`
}

type UpdateInstitutionAccreditationRequest struct {
	ID                int64   `json:"id" validate:"required,min=1"`
	AccreditationType *string `json:"accreditationType" validate:"omitempty"`
	GivenBy           *string `json:"givenBy" validate:"omitempty"`
}

// ---- RESPONSE DTOs ----

type InstitutionAccreditationResponse struct {
	ID                int64     `json:"id"`
	AccreditationType string    `json:"accreditationType"`
	GivenBy           string    `json:"givenBy"`
	CreatedAt         time.Time `json:"createdAt"`
	UpdatedAt         time.Time `json:"updatedAt"`
}
