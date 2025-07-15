package dtos

import "time"

// ---- REQUEST DTOs ----

type CreateEmployeePublicationRequest struct {
	EmployeeID        int64  `json:"employeeID" validate:"required,min=1"`
	LanguageCode      string `json:"-" validate:"required,len=2"`
	PublicationTitle  string `json:"publicationTitle" validate:"required"`
	LinkToPublication string `json:"linkToPublication" validate:"required"`
}

type UpdateEmployeePublicationRequest struct {
	ID                int64   `json:"id" validate:"required,min=1"`
	PublicationTitle  *string `json:"publicationTitle" validate:"omitempty"`
	LinkToPublication *string `json:"linkToPublication" validate:"omitempty"`
}

// ---- RESPONSE DTOs ----

type EmployeePublicationResponse struct {
	ID                int64     `json:"id"`
	PublicationTitle  string    `json:"publicationTitle"`
	LinkToPublication string    `json:"linkToPublication"`
	CreatedAt         time.Time `json:"createdAt"`
	UpdatedAt         time.Time `json:"updatedAt"`
}
