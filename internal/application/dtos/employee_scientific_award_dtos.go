package dtos

import "time"

// ---- REQUEST DTOs ----

type CreateEmployeeScientificAwardRequest struct {
	EmployeeID                int64  `json:"employeeID" validate:"required,min=1"`
	LanguageCode              string `json:"-" validate:"required,len=2"`
	ScientificAwardTitle      string `json:"scientificAwardTitle" validate:"required"`
	GivenBy                   string `json:"givenBy" validate:"required"`
}

type UpdateEmployeeScientificAwardRequest struct {
	ID                        int64   `json:"id" validate:"required,min=1"`
	ScientificAwardTitle      *string `json:"scientificAwardTitle" validate:"omitempty"`
	GivenBy                   *string `json:"givenBy" validate:"omitempty"`
}

// ---- RESPONSE DTOs ----

type EmployeeScientificAwardResponse struct {
	ID                        int64     `json:"id"`
	ScientificAwardTitle      string    `json:"scientificAwardTitle"`
	GivenBy                   string    `json:"givenBy"`
	CreatedAt                 time.Time `json:"createdAt"`
	UpdatedAt                 time.Time `json:"updatedAt"`
}
