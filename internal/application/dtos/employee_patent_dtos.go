package dtos

import "time"

// ---- REQUEST DTOs ----

type CreateEmployeePatentRequest struct {
	EmployeeID       int64  `json:"employeeID" validate:"required,min=1"`
	LanguageCode     string `json:"-" validate:"required,len=2"`
	PatentTitle      string `json:"patentTitle" validate:"required"`
	Description      string `json:"description" validate:"required"`
}

type UpdateEmployeePatentRequest struct {
	ID               int64  `json:"id" validate:"required,min=1"`
	PatentTitle      *string `json:"patentTitle" validate:"omitempty"`
	Description      *string `json:"description" validate:"omitempty"`
}

// ---- RESPONSE DTOs ----

type EmployeePatentResponse struct {
	ID               int64     `json:"id"`
	PatentTitle      string    `json:"patentTitle"`
	Description      string    `json:"description"`
	CreatedAt        time.Time `json:"createdAt"`
	UpdatedAt        time.Time `json:"updatedAt"`
}
