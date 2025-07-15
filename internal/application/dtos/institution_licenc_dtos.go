package dtos

import "time"

// ---- REQUEST DTOs ----

type CreateInstitutionLicenceRequest struct {
	InstitutionID int64     `json:"institutionID" validate:"required,min=1"`
	LanguageCode  string    `json:"-" validate:"required,len=2"`
	LicenceTitle  string    `json:"licenceTitle" validate:"required"`
	LicenceType   string    `json:"licenceType" validate:"required"`
	GivenBy       string    `json:"givenBy" validate:"required"`
	LinkToFile    string    `json:"linkToFole" validate:"required"`
	DateStart     time.Time `json:"dateStart" validate:"required"`
	DateEnd       time.Time `json:"dateEnd" validate:"required"`
}

type UpdateInstitutionLicenceRequest struct {
	ID           int64      `json:"id" validate:"required,min=1"`
	LicenceTitle *string    `json:"licenceTitle" validate:"omitempty"`
	LicenceType  *string    `json:"licenceType" validate:"omitempty"`
	GivenBy      *string    `json:"givenBy" validate:"omitempty"`
	LinkToFile   *string    `json:"linkToFole" validate:"omitempty"`
	DateStart    *time.Time `json:"dateStart" validate:"omitempty"`
	DateEnd      *time.Time `json:"dateEnd" validate:"omitempty"`
}

// ---- RESPONSE DTOs ----

type InstitutionLicenceResponse struct {
	ID            int64     `json:"id"`
	LicenceTitle  string    `json:"licenceTitle"`
	LicenceType   string    `json:"licenceType"`
	GivenBy       string    `json:"givenBy"`
	LinkToFile    string    `json:"linkToFole"`
	DateStart     time.Time `json:"dateStart"`
	DateEnd       time.Time `json:"dateEnd"`
	CreatedAt     time.Time `json:"createdAt"`
	UpdatedAt     time.Time `json:"updatedAt"`
}
