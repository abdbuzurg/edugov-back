package dtos

import "time"

// ---- REQUEST DTOs ----

type CreateEmployeeSocialRequest struct {
	EmployeeID   int64  `json:"employeeID" validate:"required,min=1"`
	LanguageCode string `json:"-" validate:"required,len=2"`
	SocialName   string `json:"socialName" validate:"required"`
	LinkToSocial string `json:"linkToSocial" validate:"required"`
}

type UpdateEmployeeSocialRequest struct {
	ID           int64   `json:"id" validate:"required,min=1"`
	SocialName   *string `json:"socialName" validate:"omitempty"`
	LinkToSocial *string `json:"linkToSocial" validate:"omitempty"`
}

// ---- RESPONSE DTOs ----

type EmployeeSocialResponse struct {
	ID           int64     `json:"id"`
	SocialName   string    `json:"socialName"`
	LinkToSocial string    `json:"linkToSocial"`
	CreatedAt    time.Time `json:"createdAt"`
	UpdatedAt    time.Time `json:"updatedAt"`
}
