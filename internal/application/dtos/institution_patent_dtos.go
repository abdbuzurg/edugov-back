package dtos

import "time"

// ---- REQUEST DTOs ----

type CreateInstitutionPatentRequest struct {
	InstitutionID     int64     `json:"institutionID" validate:"required,min=1"`
	LanguageCode      string    `json:"-" validate:"required,len=2"`
	PatentTitle       string    `json:"patentTitle" validate:"required"`
	Discipline        string    `json:"discipline" validate:"required"`
	Description       string `json:"description" validate:"required"`
	ImplementedIn     string    `json:"implementedIn" validate:"required"`
	LinkToPartnerFile string    `json:"linkToPartnerFile" validate:"required"`
}

type UpdateInstitutionPatentRequest struct {
	ID                int64      `json:"id" validate:"required,min=1"`
	PatentTitle       *string    `json:"patentTitle" validate:"omitempty"`
	Discipline        *string    `json:"discipline" validate:"omitempty"`
	Description       *string `json:"description" validate:"omitempty"`
	ImplementedIn     *string    `json:"implementedIn" validate:"omitempty"`
	LinkToPartnerFile *string    `json:"linkToPartnerFile" validate:"omitempty"`
}

// ---- RESPONSE DTOs ----

type InstitutionPatentResponse struct {
	ID                int64     `json:"id"`
	PatentTitle       string    `json:"patentTitle"`
	Discipline        string    `json:"discipline"`
	Description       string `json:"description"`
	ImplementedIn     string    `json:"implementedIn"`
	LinkToPartnerFile string    `json:"linkToPartnerFile"`
	CreatedAt         time.Time `json:"createdAt"`
	UpdatedAt         time.Time `json:"updatedAt"`
}
