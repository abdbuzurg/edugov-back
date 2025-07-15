package dtos

import "time"

// ---- REQUEST DTOs ----

type CreateInstitutionMagazineRequest struct {
	InstitutionID int64  `json:"institutionID" validate:"required,min=1"`
	LanguageCode  string `json:"-" validate:"required,len=2"`
	MagazineName  string `json:"magazineName" validate:"required"`
	Link          string `json:"link" validate:"required"`
	LinkToRINC    string `json:"linkToRINC" validate:"required"`
}

type UpdateInstitutionMagazineRequest struct {
	ID           int64   `json:"id" validate:"required,min=1"`
	MagazineName *string `json:"magazineName" validate:"omitempty"`
	Link         *string `json:"link" validate:"omitempty"`
	LinkToRINC   *string `json:"linkToRINC" validate:"omitempty"`
}

// ---- RESPONSE DTOs ----

type InstitutionMagazineResponse struct {
	ID           int64     `json:"id"`
	MagazineName string    `json:"magazineName"`
	Link         string    `json:"link"`
	LinkToRINC   string    `json:"linkToRINC"`
	CreatedAt    time.Time `json:"createdAt"`
	UpdatedAt    time.Time `json:"updatedAt"`
}
