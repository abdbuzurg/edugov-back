package dtos

import "time"

// ---- REQUEST DTOs ----

type CreateInstitutionSocialRequest struct {
	InstitutionID int64  `json:"institutionID" validate:"required,min=1"`
	LinkToSocial  string `json:"linkToSocial" validate:"required"`
	SocialName    string `json:"socialName" validate:"required"`
}

type UpdateInstitutionSocialRequest struct {
	ID           int64   `json:"id" validate:"required,min=1"`
	LinkToSocial *string `json:"linkToSocial" validate:"omitempty"`
	SocialName   *string `json:"socialName" validate:"omitempty"`
}

// ---- RESPONSE DTOs ----

type InstitutionSocialResponse struct {
	ID           int64     `json:"id"`
	LinkToSocial string    `json:"linkToSocial"`
	SocialName   string    `json:"socialName"`
	CreatedAt    time.Time `json:"createdAt"`
	UpdatedAt    time.Time `json:"updatedAt"`
}
