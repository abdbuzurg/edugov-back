package dtos

import "time"

// ---- REQUEST DTOs ----

type CreateInstitutionConferenceRequest struct {
	InstitutionID    int64     `json:"institutionID" validate:"required,min=1"`
	LanguageCode     string    `json:"-" validate:"required,len=2"`
	ConferenceTitle  string    `json:"conferenceTitle" validate:"required"`
	Link             string    `json:"link" validate:"required"`
	LinkToRINC       string    `json:"linkToRINC" validate:"required"`
	DateOfConference time.Time `json:"dateOfConference" validate:"required"`
}

type UpdateInstitutionConferenceRequest struct {
	ID               int64     `json:"id" validate:"required,min=1"`
	ConferenceTitle  *string    `json:"conferenceTitle" validate:"omitempty"`
	Link             *string    `json:"link" validate:"omitempty"`
	LinkToRINC       *string    `json:"linkToRINC" validate:"omitempty"`
	DateOfConference *time.Time `json:"dateOfConference" validate:"omitempty"`
}

// ---- RESPONSE DTOs ----

type InstitutionConferenceResponse struct {
	ID               int64     `json:"id"`
	ConferenceTitle  string    `json:"conferenceTitle"`
	Link             string    `json:"link"`
	LinkToRINC       string    `json:"linkToRINC"`
	DateOfConference time.Time `json:"dateOfConference"`
	CreatedAt        time.Time `json:"createdAt"`
	UpdatedAt        time.Time `json:"updatedAt"`
}
