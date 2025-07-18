package dtos

import "time"

// ---- REQUEST DTOs ----

type CreateEmployeeWorkExperienceRequest struct {
	EmployeeID   int64     `json:"employeeID" validate:"required,min=1"`
	LanguageCode string    `json:"-" validate:"required,len=2"`
	Workplace    string    `json:"workplace" validate:"required"`
	Description  string    `json:"description" validate:"required"`
	JobTitle     string    `json:"jobTitle" validate:"required"`
	DateStart    time.Time `json:"dateStart" validate:"required"`
	DateEnd      time.Time `json:"dateEnd" validate:"required"`
}

type UpdateEmployeeWorkExperienceRequest struct {
	ID          int64      `json:"id" validate:"required,min=1"`
	JobTitle    *string    `json:"jobTitle" validate:"omitempty"`
	Workplace   *string    `json:"workplace" validate:"omitempty"`
	Description *string    `json:"description" validate:"omitempty"`
	DateStart   *time.Time `json:"dateStart" validate:"omitempty"`
	DateEnd     *time.Time `json:"dateEnd" validate:"omitempty"`
}

// ---- RESPONSE DTOs ----

type EmployeeWorkExperienceResponse struct {
	ID            int64     `json:"id"`
	InstitutionID int64     `json:"institutionID"`
	Workplace     string    `json:"workplace"`
	Description   string    `json:"description"`
	JobTitle      string    `json:"jobTitle"`
	DateStart     time.Time `json:"dateStart"`
	DateEnd       time.Time `json:"dateEnd"`
	CreatedAt     time.Time `json:"createdAt"`
	UpdatedAt     time.Time `json:"updatedAt"`
}
