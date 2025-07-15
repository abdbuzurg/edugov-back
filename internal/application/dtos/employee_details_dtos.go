package dtos

import "time"

// ---- REQUEST DTOS ----

type CreateEmployeeDetailsRequest struct {
	LanguageCode         string `json:"languageCode" validate:"required,len=2"`
	Surname              string `json:"surname" validate:"required"`
	Name                 string `json:"name" validate:"required"`
	Middlename           string `json:"middlename" validate:"required"`
	IsEmployeeDetailsNew bool   `json:"isNewEmployeeDetailsNew" validate:"required"`
}

type UpdateEmployeeDetailsRequest struct {
	ID                   int64   `json:"id" validate:"min=0"`
	EmployeeID           int64   `json:"employeeID" validate:"required,min=1"`
	Surname              *string `json:"surname"`
	Name                 *string `json:"name"`
	Middlename           *string `json:"middlename"`
	IsEmployeeDetailsNew *bool   `json:"isNewEmployeeDetailsNew"`
}

type UpdateFullEmployeeData struct {
	Data []UpdateEmployeeDetailsRequest `json:"data"`
}

// ---- RESPONSE DTOS ----

type EmployeeDetailsResponse struct {
	ID                   int64     `json:"id"`
	Surname              string    `json:"surname"`
	Name                 string    `json:"name"`
	Middlename           string    `json:"middlename"`
	IsEmployeeDetailsNew bool      `json:"isNewEmployeeDetails"`
	CreatedAt            time.Time `json:"createdAt"`
	UpdatedAt            time.Time `json:"updatedAt"`
}
