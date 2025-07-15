package dtos

import "time"

// ---- REQUEST DTOs ----

type CreateEmployeeDegreeRequest struct {
	EmployeeID         int64     `json:"employeeID" validate:"required,min=1"`
	LanguageCode       string    `json:"-" validate:"required,len=2"`
	DegreeLevel        string    `json:"degreeLevel" validate:"required"`
	UniversityName     string    `json:"universityName" validate:"required"`
	Speciality         string    `json:"speciality" validate:"required"`
	DateStart          time.Time `json:"dateStart" validate:"required"`
	DateEnd            time.Time `json:"dateEnd" validate:"required"`
	GivenBy            string    `json:"givenBy" validate:"required"`
	DateDegreeRecieved time.Time `json:"dateDegreeRecieved" validate:"required"`
	LinkToDegreeFile   string    `json:"linkToDegreeFile" validate:"required"`
}

type UpdateEmployeeDegreeRequest struct {
	ID                 int64      `json:"id" validate:"required,min=1"`
	DegreeLevel        *string    `json:"degreeLevel" validate:"omitempty"`
	UniversityName     *string    `json:"universityName" validate:"omitempty"`
	Speciality         *string    `json:"speciality" validate:"omitempty"`
	DateStart          *time.Time `json:"dateStart" validate:"omitempty"`
	DateEnd            *time.Time `json:"dateEnd" validate:"omitempty"`
	GivenBy            *string    `json:"givenBy" validate:"omitempty"`
	DateDegreeRecieved *time.Time `json:"dateDegreeRecieved" validate:"omitempty"`
	LinkToDegreeFile   *string    `json:"linkToDegreeFile" validate:"omitempty"`
}

// ---- RESPONSE DTOs ----

type EmployeeDegreeResponse struct {
	ID                 int64     `json:"id"`
	DegreeLevel        string    `json:"degreeLevel"`
	UniversityName     string    `json:"universityName"`
	Speciality         string    `json:"speciality"`
	DateStart          time.Time `json:"dateStart"`
	DateEnd            time.Time `json:"dateEnd"`
	GivenBy            string    `json:"givenBy"`
	DateDegreeRecieved time.Time `json:"dateDegreeRecieved"`
	LinkToDegreeFile   string    `json:"linkToDegreeFile"`
	CreatedAt          time.Time `json:"createdAt"`
	UpdatedAt          time.Time `json:"updatedAt"`
}
