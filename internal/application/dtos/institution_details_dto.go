package dtos

import "time"

// ---- REQUEST DTOs ----

type CreateInstitutionDetailsRequest struct {
	LanguageCode     string  `json:"languageCode" validate:"required,len=2"`
	InstitutionTitle string  `json:"institutionTitle" validate:"required"`
	InstitutionType  string  `json:"institutionType" validate:"required"`
	LegalStatus      string  `json:"legalStatus" validate:"required"`
	Mission          string  `json:"mission" validate:"required"`
	Founder          string  `json:"founder" validate:"required"`
	LegalAddress     string  `json:"legalAddress" validate:"required"`
	FactualAddress   *string `json:"factualAddress"`
}

type UpdateInstitutionDetailsRequest struct {
	ID               int64   `json:"id" validate:"required,min=1"`
	LanguageCode     string  `json:"languageCode" validate:"required,len=2"`
	InstitutionTitle *string `json:"institutionTitle,omitempty"`
	InstitutionType  *string `json:"institutionType,omitempty"`
	LegalStatus      *string `json:"legalStatus,omitempty"`
	Mission          *string `json:"mission,omitempty"`
	Founder          *string `json:"founder,omitempty"`
	LegalAddress     *string `json:"legalAddress,omitempty"`
	FactualAddress   *string `json:"factualAddress,omitempty"`
}

// ---- RESPONSE DTOs ----

type InstitutionDetailsResponse struct {
	ID               int64    `json:"id"`
	InstitutionTitle string    `json:"institutionTitle"`
	InstitutionType  string    `json:"institutionType"`
	LegalStatus      string    `json:"legalStatus"`
	Mission          string    `json:"mission"`
	Founder          string    `json:"founder"`
	LegalAddress     string    `json:"legalAddress"`
	FactualAddress   *string   `json:"factualAddress,omitempty"`
	CreatedAt        time.Time `json:"createdAt"`
	UpdatedAt        time.Time `json:"updatedAt"`
}
